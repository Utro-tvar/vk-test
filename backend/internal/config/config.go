package config

import (
	"fmt"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func FromEnv() (*Config, error) {
	const op = "config.FromEnv"

	cfg := Config{}
	var exists bool

	cfg.Host, exists = os.LookupEnv("DB_HOST")
	if !exists {
		return nil, fmt.Errorf("%s: DB_HOST env var not found", op)
	}
	cfg.Port, exists = os.LookupEnv("DB_PORT")
	if !exists {
		return nil, fmt.Errorf("%s: DB_PORT env var not found", op)
	}
	cfg.User, exists = os.LookupEnv("DB_USER")
	if !exists {
		return nil, fmt.Errorf("%s: DB_USER env var not found", op)
	}
	cfg.Password, exists = os.LookupEnv("DB_PASS")
	if !exists {
		return nil, fmt.Errorf("%s: DB_PASS env var not found", op)
	}
	cfg.DBName, exists = os.LookupEnv("DB_NAME")
	if !exists {
		return nil, fmt.Errorf("%s: DB_NAME env var not found", op)
	}

	return &cfg, nil
}
