package main

import (
	"context"
	"log"
	"os"

	"github.com/kngnkg/tunetrail/restapi/config"
	"github.com/kngnkg/tunetrail/restapi/router"
	"github.com/kngnkg/tunetrail/restapi/runner"
	"github.com/kngnkg/tunetrail/restapi/validate"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Printf("cannot get config: %v", err)
		os.Exit(1)
	}

	// バリデーションの初期化
	if err := validate.InitValidation(); err != nil {
		log.Printf("cannot init validation: %v", err)
		os.Exit(1)
	}

	// ルーターの初期化
	r, cleanup, err := router.SetupRouter(cfg)
	if err != nil {
		log.Printf("cannot setup router: %v", err)
		os.Exit(1)
	}
	defer cleanup()

	// サーバーの起動
	if err := runner.Run(context.Background(), r, cfg); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}
