package database

import (
	"log"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	PostgresURL string
	RedisURL    string
}

type Store struct {
	Config Config
	DB     *gorm.DB
	Cache  *redis.Client
}

func New(config Config) *Store {
	DB, err := gorm.Open(postgres.Open(config.PostgresURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	Cache := redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})

	return &Store{
		Config: config,
		DB:     DB,
		Cache:  Cache,
	}
}
