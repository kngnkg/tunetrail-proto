package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kwtryo/tunetrail/database/config"
	"github.com/kwtryo/tunetrail/database/s3downloader"
)

type Event struct {
	Item string `json:"item"` // バケット内のオブジェクトのキー
}

type Response struct {
	Result string `json:"result"` // レスポンスメッセージ
}

func handleRequest(ctx context.Context, event Event) (Response, error) {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("your-region"),
	})
	if err != nil {
		log.Fatalf("Error creating AWS session: %v", err)
	}

	s3dl := s3downloader.New(sess, cfg.S3Bucket)
	file, err := s3dl.Download(ctx, event.Item)
	if err != nil {
		log.Fatalf("Error downloading from S3: %v", err)
	}

	if err := migration(cfg, file); err != nil {
		log.Fatalf("Error running migration: %v", err)
	}

	return Response{Result: "Success"}, nil
}

func main() {
	if os.Getenv("ENV") == "dev" {
		log.Println("Running locally")
		handleRequest(context.Background(), Event{Item: "local.sql"})
		return
	}

	log.Println("Starting Lambda")
	lambda.Start(handleRequest)
}
