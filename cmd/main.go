package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/chat"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/config"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/server"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/admin"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/auth"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/category"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/citizencard"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/customer"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/media"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/message"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/objects"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/packages"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/photographers"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/quotation"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/report"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/review"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/stripe"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/user"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/validator"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
)

// @title						Pic Me Pls API
// @version						0.1
// @description					Pic Me Pls API Documentation
// @securityDefinitions.apikey ApiKeyAuth
// @in							header
// @name						Authorization
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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
	chatService := chat.NewServer(store, validate)

	// middlewares
	authMiddleware := authentication.NewAuthMiddleware(jwtService)

	// handlers
	authHandler := auth.NewHandler(store, validate, jwtService, authMiddleware, config.GoogleClientID)
	userHandler := user.NewHandler(store, validate, authMiddleware)
	photographersHandler := photographers.NewHandler(store, validate, authMiddleware)
	citizencardHandler := citizencard.NewHandler(store, validate, authMiddleware)
	packageHandler := packages.NewHandler(store, validate, authMiddleware)
	reviewHandler := review.NewHandler(store, validate, authMiddleware)
	categoryHandler := category.NewHandler(store, validate)
	objectHandler := objects.NewHandler(store, config.Storage)
	quotationHandler := quotation.NewHandler(store, authMiddleware, validate, chatService)
	mediaHandler := media.NewHandler(store, validate, authMiddleware)
	customerHandler := customer.NewHandler(store, validate)
	messageHandler := message.NewHandler(store, authMiddleware, chatService)
	stripeHandler := stripe.NewHandler(store, validate, authMiddleware, config.Stripe, config.FrontendURL)
	adminHandler := admin.NewHandler(store, authMiddleware, validate)
	reportHandler := report.NewHandler(store, authMiddleware, validate)

	server.RegisterDocs()

	// routes
	server.RegisterRoutes(
		authMiddleware,
		authHandler,
		userHandler,
		photographersHandler,
		citizencardHandler,
		packageHandler,
		reviewHandler,
		categoryHandler,
		objectHandler,
		quotationHandler,
		mediaHandler,
		customerHandler,
		messageHandler,
		stripeHandler,
		adminHandler,
		reportHandler,
	)

	server.Start(ctx, stop)
}
