package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/kwtryo/tunetrail/database/config"
)

func migration(cfg *config.Config, file *os.File) error {
	cmd := exec.Command("psqldef",
		"--dry-run",
		cfg.DBName,
		"--host="+cfg.DBHost,
		"--port="+strconv.Itoa(cfg.DBPort),
		"--user="+cfg.DBUser,
		"--password="+cfg.DBPassword,
		"--file="+file.Name(),
	)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	outStr, errStr := stdout.String(), stderr.String()
	log.Printf("stdout: %s", outStr)
	if err != nil {
		log.Printf("stderr: %s", errStr)
		log.Fatalf("psqldef failed with %s\n", err)
	} else if stderr.Len() > 0 {
		// 警告が出ても処理は続行する
		log.Printf("stderr: %s", errStr)
	}

	return nil
}
