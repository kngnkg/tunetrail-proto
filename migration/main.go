package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/postgres"
	"github.com/k0kubun/sqldef/schema"
	"github.com/kngnkg/tunetrail/database/config"
	"github.com/kngnkg/tunetrail/database/s3downloader"
)

type Event struct {
	Items []string `json:"items"` // バケット内のオブジェクトのキー
}

type Response struct {
	Result string `json:"result"` // レスポンスメッセージ
}

func getSchemaFilesFromS3(ctx context.Context, region string, S3Bucket string, objectKeys []string) ([]string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	s3dl := s3downloader.New(sess, S3Bucket)
	var files []string
	for _, objectKey := range objectKeys {
		file, err := s3dl.Download(ctx, objectKey)
		if err != nil {
			return nil, err
		}
		files = append(files, file.Name())
	}
	return files, nil
}

func migrate(ctx context.Context, cfg *config.Config, files []string) error {
	desiredFiles := sqldef.ParseFiles(files)
	desiredDDLs, err := sqldef.ReadFiles(desiredFiles)
	if err != nil {
		return err
	}

	options := &sqldef.Options{
		DesiredDDLs:     desiredDDLs,
		DryRun:          cfg.DryRun,
		EnableDropTable: cfg.EnableDropTable,
	}

	dbCfg := database.Config{
		DbName:   cfg.DBName,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
	}

	db, err := postgres.NewDatabase(dbCfg)
	if err != nil {
		return err
	}

	sqlParser := postgres.NewParser()

	if cfg.Env == "dev" {
		os.Setenv("PGSSLMODE", "disable")
	}
	sqldef.Run(schema.GeneratorModePostgres, db, sqlParser, options)
	return nil
}

func handleRequest(ctx context.Context, event Event) (Response, error) {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	files, err := getSchemaFilesFromS3(ctx, cfg.AWSRegion, cfg.S3Bucket, event.Items)
	if err != nil {
		log.Fatalf("Error getting schema file: %v", err)
	}

	if err := migrate(ctx, cfg, files); err != nil {
		log.Fatalf("Error migration: %v", err)
	}

	return Response{Result: "Success"}, nil
}

func main() {
	log.Println("Starting Lambda")
	lambda.Start(handleRequest)
}
