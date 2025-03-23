package database

import (
	"context"
	"log"
	"log/slog"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	pglib "github.com/CP-RektMart/pic-me-pls-backend/pkg/postgres"
	rdlib "github.com/CP-RektMart/pic-me-pls-backend/pkg/redis"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/storage"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB      *gorm.DB
	Cache   *redis.Client
	Storage *storage.Client
}

func New(ctx context.Context, pgConfig pglib.Config, rdConfig rdlib.Config, storageConfig storage.Config) *Store {
	db, err := gorm.Open(postgres.Open(pgConfig.String()), &gorm.Config{})
	if err != nil {
		logger.PanicContext(ctx, "failed to connect to database", slog.Any("error", err))
	}

	redisConn, err := rdlib.New(ctx, rdConfig)
	if err != nil {
		logger.PanicContext(ctx, "failed to connect to redis", slog.Any("error", err))
	}

	storage, err := storage.New(ctx, storageConfig)
	if err != nil {
		logger.PanicContext(ctx, "failed to connect to storage", slog.Any("error", err))
	}

	store := &Store{
		DB:      db,
		Cache:   redisConn,
		Storage: storage,
	}
	store.migrate()
	return store
}

func (s *Store) migrate() {
	log.Println("Running migrations...")

	if err := s.DB.AutoMigrate(
		&model.User{},
		&model.Photographer{},
		&model.Package{},
		&model.Tag{},
		&model.Category{},
		&model.Media{},
		&model.Message{},
		&model.Quotation{},
		&model.Review{},
		&model.CitizenCard{},
		&model.Preview{},
	); err != nil {
		logger.Panic("failed to migrate database", slog.Any("error", err))
	}

	log.Println("Migrations complete!")
}
