package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/requestlogger"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
	db     *database.Store
}

func New(config Config, corsConfig CorsConfig, db *database.Store) *Server {
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

	return &Server{
		config: config,
		app:    app,
		db:     db,
	}
}

func (s *Server) Start(ctx context.Context, stop context.CancelFunc) {
	// Health check
	s.app.Get("/v1/", func(c *fiber.Ctx) error {
		return c.JSON(dto.HttpResponse{
			Result: "ok",
		})
	})

	// Example upload file
	s.app.Post("/v1/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			apperror.BadRequest("failed to get file", err)
		}

		contentType := file.Header.Get("Content-Type")

		src, err := file.Open()
		if err != nil {
			return errors.Wrap(err, "failed to open file")
		}
		defer src.Close()

		if err := s.db.Storage.UploadFile(ctx, file.Filename, contentType, src, true); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.HttpResponse{
				Error: err.Error(),
			})
		}

		return c.JSON(dto.HttpResponse{
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
