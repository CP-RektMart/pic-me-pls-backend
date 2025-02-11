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
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/requestlogger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
)

type Config struct {
	Name         string `env:"NAME"`
	Port         int    `env:"PORT"`
	MaxBodyLimit int    `env:"MAX_BODY_LIMIT"`
}

type CorsConfig struct {
	AllowedOrigins   string `env:"ALLOWED_ORIGINS"`
	AllowedMethods   string `env:"ALLOWED_METHODS"`
	AllowedHeaders   string `env:"ALLOWED_HEADERS"`
	AllowCredentials bool   `env:"ALLOW_CREDENTIALS"`
}

type Server struct {
	config Config
	app    *fiber.App
}

func New(config Config, corsConfig CorsConfig, jwtConfig jwt.Config, db *database.Store) *Server {
	app := fiber.New(fiber.Config{
		AppName:       config.Name,
		BodyLimit:     config.MaxBodyLimit * 1024 * 1024,
		CaseSensitive: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return apperror.Internal("internal server error", err)
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

	return &Server{
		config: config,
		app:    app,
	}
}

func (s *Server) Start(ctx context.Context, stop context.CancelFunc) {
	s.app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(dto.HttpResponse[string]{
			Result: "ok",
		})
	})

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
