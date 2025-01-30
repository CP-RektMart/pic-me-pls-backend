package logger

import (
	"context"
	"log/slog"
	"os"
)

type Config struct {
	Debug  bool   `env:"DEBUG"`
	Output string `env:"OUTPUT"`
}

var (
	logger *slog.Logger

	LevelPanic = slog.Level(14)
)

func Init(config Config) error {
	loggerLevel := new(slog.LevelVar)

	if config.Debug {
		loggerLevel.Set(slog.LevelDebug)
	}

	switch config.Output {
	case "json":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	default:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}

	return nil
}

func log(ctx context.Context, level slog.Level, msg string, args ...any) {
	logger.Log(ctx, level, msg, args...)
}

func Info(msg string, args ...any) {
	log(context.Background(), slog.LevelInfo, msg, args...)
}

func Debug(msg string, args ...any) {
	log(context.Background(), slog.LevelDebug, msg, args...)
}

func Error(msg string, args ...any) {
	log(context.Background(), slog.LevelError, msg, args...)
}

func Warn(msg string, args ...any) {
	log(context.Background(), slog.LevelWarn, msg, args...)
}

func Panic(msg string, args ...any) {
	log(context.Background(), LevelPanic, msg, args...)
	panic(msg)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelInfo, msg, args...)
}

func DebugContext(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelDebug, msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelError, msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	log(ctx, slog.LevelWarn, msg, args...)
}

func PanicContext(ctx context.Context, msg string, args ...any) {
	log(ctx, LevelPanic, msg, args...)
	panic(msg)
}
