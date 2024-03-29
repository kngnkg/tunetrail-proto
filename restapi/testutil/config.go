package testutil

import (
	"os"
	"testing"

	"github.com/caarlos0/env"
	"github.com/kngnkg/tunetrail/restapi/config"
)

// テスト用のconfigを返す
func CreateConfigForTest(t *testing.T) *config.Config {
	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		t.Fatalf("cannot parse env: %v", err)
	}
	cfg.Port = 8081
	if _, defined := os.LookupEnv("CI"); defined {
		cfg.DBHost = "127.0.0.1"
	}
	return cfg
}
