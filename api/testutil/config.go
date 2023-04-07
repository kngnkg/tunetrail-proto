package testutil

import (
	"os"
	"testing"

	"github.com/caarlos0/env"
	"github.com/kwtryo/tunetrail/api/config"
)

// テスト用のconfigを返す
func CreateConfigForTest(t *testing.T) *config.Config {
	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		t.Fatalf("cannot parse env: %v", err)
	}
	cfg.Port = 8081
	// CI環境ならポート番号を変更
	if _, defined := os.LookupEnv("CI"); defined {
		cfg.DBPort = 5432
	}
	return cfg
}
