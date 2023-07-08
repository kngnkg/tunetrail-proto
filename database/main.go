package main

// type Config struct {
// 	DBHost     string `env:"TUNETRAIL_DB_HOST" envDefault:"tunetrail-db"`
// 	DBPort     int    `env:"TUNETRAIL_DB_PORT" envDefault:"5432"`
// 	DBUser     string `env:"TUNETRAIL_DB_USER" envDefault:"tunetrail"`
// 	DBPassword string `env:"TUNETRAIL_DB_PASSWORD" envDefault:"tunetrail"`
// 	DBName     string `env:"TUNETRAIL_DB_NAME" envDefault:"tunetrail"`
// }

// func loadConfig() (*Config, error) {
// 	cfg := &Config{}
// 	if err := env.Parse(cfg); err != nil {
// 		return nil, err
// 	}
// 	return cfg, nil
// }

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	log.Printf("Processing Lambda request %s\n", name.Name)
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	log.Printf("Starting Lambda\n")
	lambda.Start(HandleRequest)
}
