package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env        string `env:"TUNETRAIL_ENV" envDefault:"dev"`
	S3Bucket   string `env:"TUNETRAIL_S3_BUCKET" envDefault:"tunetrail"`
	DBHost     string `env:"TUNETRAIL_DB_HOST" envDefault:"tunetrail-db"`
	DBPort     int    `env:"TUNETRAIL_DB_PORT" envDefault:"5432"`
	DBUser     string `env:"TUNETRAIL_DB_USER" envDefault:"tunetrail"`
	DBPassword string `env:"TUNETRAIL_DB_PASSWORD" envDefault:"tunetrail"`
	DBName     string `env:"TUNETRAIL_DB_NAME" envDefault:"tunetrail"`
	DryRun     bool   `env:"DRY_RUN" envDefault:"true"`

	// 以下はAWS Lambdaの環境変数
	AWSRegion string `env:"AWS_REGION" envDefault:"ap-northeast-1"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
