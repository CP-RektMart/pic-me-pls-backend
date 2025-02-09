package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/requestlogger"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Config struct {
	Name               string `env:"NAME"`
	Port               int    `env:"PORT"`
	MaxBodyLimit       int    `env:"MAX_BODY_LIMIT"`
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	JwtAccessSecret    string `env:"JWT_ACCESS_SECRET"`
	JWTRefreshSecret   string `env:"JWT_REFRESH_SECRET"`
	JwtAccessDuration  int    `env:"JWT_ACCESS_DURATION"`
	JwtRefreshDuration int    `env:"JWT_REFRESH_DURATION"`
}

type CorsConfig struct {
	AllowedOrigins   string `env:"ALLOWED_ORIGINS"`
	AllowedMethods   string `env:"ALLOWED_METHODS"`
	AllowedHeaders   string `env:"ALLOWED_HEADERS"`
	AllowCredentials bool   `env:"ALLOW_CREDENTIALS"`
}

type Server struct {
	config         Config
	app            *fiber.App
	db             *database.Store
	authMiddleware authentication.AuthMiddleware
}

func New(config Config, corsConfig CorsConfig, db *database.Store, authMiddleware authentication.AuthMiddleware) *Server {
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

	return &Server{
		config:         config,
		app:            app,
		db:             db,
		authMiddleware: authMiddleware,
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

	// Example upload
	v1.Post("/upload", func(c *fiber.Ctx) error {
		ctx := c.UserContext()

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
			return errors.Wrap(err, "failed to upload file")
		}

		return c.JSON(dto.HttpResponse{
			Result: "ok",
		})
	})
}
