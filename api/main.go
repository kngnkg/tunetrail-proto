package main

import (
	"context"
	"log"
	"os"

	"github.com/kwtryo/tunetrail/api/config"
	"github.com/kwtryo/tunetrail/api/router"
	"github.com/kwtryo/tunetrail/api/runner"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("cannot get config")
	}

	r, cleanup, err := router.SetupRouter(cfg)
	if err != nil {
		log.Printf("cannot setup router: %v", err)
		os.Exit(1)
	}
	defer cleanup()

	if err := runner.Run(context.Background(), r, cfg); err != nil {
		log.Printf("failed to terminated server: %v", err)
		os.Exit(1)
	}
}
