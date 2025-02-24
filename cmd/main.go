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
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/media"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/message"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/object"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/packages"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/photographers"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/quotation"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/review"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/user"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/verify"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/validator"
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
	server := server.New(config.Server, config.Cors, config.JWT, store)
	validate := validator.New()

	// services
	jwtService := jwt.New(config.JWT, store.Cache)

	// middlewares
	authMiddleware := authentication.NewAuthMiddleware(server.Api, jwtService)

	// handlers
	authHandler := auth.NewHandler(store, validate, jwtService, authMiddleware, config.GoogleClientID)
	userHandler := user.NewHandler(store, validate, authMiddleware)
	verifyHandler := verify.NewHandler(store, validate, authMiddleware)
	photographersHandler := photographers.NewHandler(store, validate, authMiddleware)
	packageHandler := packages.NewHandler(store, validate, authMiddleware)
	reviewHandler := review.NewHandler(store, validate)
	categoryHandler := category.NewHandler(store, validate)
	messageHandler := message.NewHandler(store, validate)
	objectHandler := object.NewHandler(store, config.Storage)
	quotationHandler := quotation.NewHandler(store, authMiddleware)
	mediaHandler := media.NewHandler(store, validate, authMiddleware)

	server.RegisterDocs()

	// routes
	server.RegisterRoutes(
		authMiddleware,
		authHandler,
		userHandler,
		verifyHandler,
		photographersHandler,
		packageHandler,
		reviewHandler,
		categoryHandler,
		messageHandler,
		objectHandler,
		quotationHandler,
		mediaHandler,
	)

	server.Start(ctx, stop)
}
