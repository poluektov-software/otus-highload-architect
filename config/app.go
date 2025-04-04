package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
	HostPort    string
}

func Get() (*Config, error) {
	var cfg Config

	hostPort := os.Getenv("HOST_PORT")
	if hostPort == "" {
		log.Fatal("В окружении не указан HOST_PORT")
	}
	cfg.HostPort = hostPort

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("В окружении не указан DATABASE_URL")
	}
	cfg.DatabaseURL = dbURL

	return &cfg, nil
}
