package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kwtryo/tunetrail/database/config"
	"github.com/kwtryo/tunetrail/database/s3dl"
)

type Event struct {
	Item string `json:"name"`
}
type Response struct {
	Message string `json:"Answer:"`
}

func handleRequest(ctx context.Context, event Event) (Response, error) {
	log.Printf("Processing Lambda request %s", event.Item)
	// item := "sample.sql"
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	log.Printf("Config: %+v", cfg)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("your-region"),
	})
	if err != nil {
		log.Fatalf("Error creating AWS session: %v", err)
	}

	downloader := s3dl.New(sess, cfg.S3Bucket)
	file, err := downloader.Download(ctx, event.Item)
	if err != nil {
		log.Fatalf("Error downloading from S3: %v", err)
	}

	if err := migration(cfg, file); err != nil {
		log.Fatalf("Error running migration: %v", err)
	}

	return Response{Message: "Success"}, nil
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
