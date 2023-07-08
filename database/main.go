package main

import (
	"log"
)

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

func main() {
	// cfg, err := loadConfig()
	// if err != nil {
	// 	log.Fatalf("Failed to parse config: %s\n", err)
	// }

	// cmd := exec.Command("psqldef",
	// 	cfg.DBName,
	// 	"--host="+cfg.DBHost,
	// 	"--port="+strconv.Itoa(cfg.DBPort),
	// 	"--user="+cfg.DBUser,
	// 	"--password="+cfg.DBPassword,
	// 	"--file=_tools/postgres/schema.sql",
	// )

	// var stdout, stderr bytes.Buffer
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr

	// err = cmd.Run()
	// outStr, errStr := stdout.String(), stderr.String()
	// log.Printf("stdout: %s\n", outStr)
	// if err != nil {
	// 	log.Printf("stderr: %s\n", errStr)
	// 	log.Fatalf("psqldef failed with %s\n", err)
	// } else if stderr.Len() > 0 {
	// 	log.Printf("stderr: %s\n", errStr)
	// }

	log.Print("hello world")
}
