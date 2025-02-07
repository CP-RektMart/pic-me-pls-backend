package database

import (
	"context"
	"log"
	"log/slog"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	pglib "github.com/CP-RektMart/pic-me-pls-backend/pkg/postgres"
	rdlib "github.com/CP-RektMart/pic-me-pls-backend/pkg/redis"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func New(ctx context.Context, pgconfig pglib.Config, rdconfig rdlib.Config) *Store {
	db, err := gorm.Open(postgres.Open(pgconfig.String()), &gorm.Config{})
	if err != nil {
		logger.PanicContext(ctx, "failed to connect to database", slog.Any("error", err))
	}

	redisConn, err := rdlib.New(ctx, rdconfig)
	if err != nil {
		logger.PanicContext(ctx, "failed to connect to redis", slog.Any("error", err))
	}
	defer func() {
		if err := redisConn.Close(); err != nil {
			logger.ErrorContext(ctx, "failed to close redis connection", slog.Any("error", err))
		}
	}()

	store := &Store{
		DB:    db,
		Cache: redisConn,
	}
	store.migrate()
	return store
}

func (s *Store) migrate() {
	log.Println("Running migrations...")
	if err := s.DB.AutoMigrate(
		&model.User{},
		&model.Photographer{},
		&model.Gallery{},
		&model.Tag{},
		&model.Category{},
		&model.Media{},
		&model.Message{},
		&model.Quotation{},
		&model.Review{},
	); err != nil {
		logger.Panic("failed to migrate database", slog.Any("error", err))
	}
	log.Println("Migrations complete!")
}
