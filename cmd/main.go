package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/config"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/server"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
)

func main() {
	// hello
	config := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := logger.Init(config.Logger); err != nil {
		logger.PanicContext(ctx, "failed to initialize logger", slog.Any("error", err))
	}

	store := database.New(ctx, config.Postgres, config.Redis, config.Storage)
	server := server.New(config.Server, config.Cors, store)

	server.Start(ctx, stop)
}
