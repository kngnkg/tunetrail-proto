package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env        string `env:"TUNETRAIL_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"80"`
	DBHost     string `env:"TUNETRAIL_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"TUNETRAIL_DB_PORT" envDefault:"5555"`
	DBUser     string `env:"TUNETRAIL_DB_USER" envDefault:"tunetrail"`
	DBPassword string `env:"TUNETRAIL_DB_PASSWORD" envDefault:"tunetrail"`
	DBName     string `env:"TUNETRAIL_DB_NAME" envDefault:"tunetrail"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
