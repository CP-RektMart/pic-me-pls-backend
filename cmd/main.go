package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/config"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/server"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/auth"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/category"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/example"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/gallery"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/message"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/photographer"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/review"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/user"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/validator"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
)

// @title pic-me-pls API
// @version 1.0
// @description pic-me-pls API documentation

// @schemes https http

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// hello
	config := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := logger.Init(config.Logger); err != nil {
		logger.PanicContext(ctx, "failed to initialize logger", slog.Any("error", err))
	}

	store := database.New(ctx, config.Postgres, config.Redis, config.Storage)
	server := server.New(config.Server, config.Cors, config.JWT, store)
	validate := validator.New()

	// services
	jwtService := jwt.New(config.JWT, store.Cache)

	// middlewares
	authMiddleware := authentication.NewAuthMiddleware(jwtService)

	// handlers
	exampleHandler := example.NewHandler(store)
	authHandler := auth.NewHandler(store, validate, jwtService, config.GoogleClientID)
	userHandler := user.NewHandler(store, validate, authMiddleware)
	photographerHandler := photographer.NewHandler(store, validate)
	galleryHandler := gallery.NewHandler(store, validate)
	reviewHandler := review.NewHandler(store, validate)
	categoryHandler := category.NewHandler(store, validate)
	messageHandler := message.NewHandler(store, validate)

	server.RegisterRoutes(
		authMiddleware,
		exampleHandler,
		authHandler,
		userHandler,
		photographerHandler,
		galleryHandler,
		reviewHandler,
		categoryHandler,
		messageHandler,
	)

	server.Start(ctx, stop)
}
