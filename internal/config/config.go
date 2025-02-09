package config

import (
	"log"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/server"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/postgres"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/redis"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/storage"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server   server.Config     `envPrefix:"SERVER_"`
	Logger   logger.Config     `envPrefix:"LOGGER_"`
	Postgres postgres.Config   `envPrefix:"POSTGRES_"`
	Redis    redis.Config      `envPrefix:"REDIS_"`
	Cors     server.CorsConfig `envPrefix:"CORS_"`
	Storage  storage.Config    `envPrefix:"STORAGE_"`
}

func Load() *AppConfig {
	appConfig := &AppConfig{}
	_ = godotenv.Load()

	if err := env.Parse(appConfig); err != nil {
		log.Fatalf("failed parse env: %s", err)
	}

	return appConfig
}
