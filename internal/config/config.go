package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	PostgresURL string `env:"POSTGRES_URL"`
	ServerAddr  string `env:"SERVER_ADDR"`
}

func Load() *AppConfig {
	appConfig := &AppConfig{}
	_ = godotenv.Load()

	if err := env.Parse(appConfig); err != nil {
		log.Fatalf("failed parse env: %s", err)
	}

	return appConfig
}
