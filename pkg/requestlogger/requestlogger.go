package requestlogger

import (
	"log/slog"

	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func New() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		requestId := ctx.GetRespHeader("X-Request-Id")
		logger.InfoContext(ctx.UserContext(), "request received", slog.String("request_id", requestId), slog.String("method", ctx.Method()), slog.String("path", ctx.Path()))
		return ctx.Next()
	}
}
