package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	_ "github.com/CP-RektMart/pic-me-pls-backend/doc"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	custom_validator "github.com/CP-RektMart/pic-me-pls-backend/internal/validator"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/requestlogger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
)

type Config struct {
	Name           string `env:"NAME"`
	Port           int    `env:"PORT"`
	MaxBodyLimit   int    `env:"MAX_BODY_LIMIT"`
	GoogleClientID string `env:"GOOGLE_CLIENT_ID"`
	JWT            jwt.Config
}

type CorsConfig struct {
	AllowedOrigins   string `env:"ALLOWED_ORIGINS"`
	AllowedMethods   string `env:"ALLOWED_METHODS"`
	AllowedHeaders   string `env:"ALLOWED_HEADERS"`
	AllowCredentials bool   `env:"ALLOW_CREDENTIALS"`
}

type Server struct {
	config     Config
	app        *fiber.App
	db         *database.Store
	validate   *validator.Validate
	middleware authentication.AuthMiddleware
}

func New(config Config, corsConfig CorsConfig, jwtConfig jwt.Config, db *database.Store) *Server {
	app := fiber.New(fiber.Config{
		AppName:       config.Name,
		BodyLimit:     config.MaxBodyLimit * 1024 * 1024,
		CaseSensitive: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			logger.ErrorContext(c.UserContext(), "unhandled error", slog.Any("error", err))
			return c.SendStatus(fiber.StatusInternalServerError)
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     corsConfig.AllowedOrigins,
		AllowMethods:     corsConfig.AllowedMethods,
		AllowHeaders:     corsConfig.AllowedHeaders,
		AllowCredentials: corsConfig.AllowCredentials,
	})).
		Use(requestid.New()).
		Use(requestlogger.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	config.JWT = jwtConfig

	return &Server{
		config:     config,
		app:        app,
		db:         db,
		validate:   custom_validator.New(),
		middleware: authentication.NewAuthMiddleware(&jwtConfig, db.Cache),
	}
}

func (s *Server) Start(ctx context.Context, stop context.CancelFunc) {
	s.app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(dto.HttpResponse{
			Result: "ok",
		})
	})

	s.registerRoute()

	go func() {
		if err := s.app.Listen(fmt.Sprintf(":%d", s.config.Port)); err != nil {
			logger.PanicContext(ctx, "failed to start server", slog.Any("error", err))
			stop()
		}
	}()

	defer func() {
		if err := s.app.ShutdownWithContext(ctx); err != nil {
			logger.ErrorContext(ctx, "failed to shutdown server", slog.Any("error", err))
		}
		logger.InfoContext(ctx, "gracefully shutdown server")
	}()

	<-ctx.Done()
	logger.InfoContext(ctx, "Shutting down server")
}

func (s *Server) registerRoute() {
	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/login", s.handleLogin)
}
