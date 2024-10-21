package logger

import (
	"chat/internal/lib/slogx"
	"fmt"
	"log/slog"
	"os"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDev:
		log = slog.New(
			slogx.NewPrettyHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
		log.Warn(fmt.Sprintf("unknown environment, used: %s", envProd))
	}

	return log
}
