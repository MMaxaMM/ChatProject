package logger

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		opts := &tint.Options{Level: slog.LevelDebug, TimeFormat: time.Kitchen}
		log = slog.New(tint.NewHandler(os.Stdout, opts))

	case envProd:
		opts := &slog.HandlerOptions{Level: slog.LevelInfo}
		log = slog.New(slog.NewJSONHandler(os.Stdout, opts))

	default:
		opts := &slog.HandlerOptions{Level: slog.LevelInfo}
		log = slog.New(slog.NewJSONHandler(os.Stdout, opts))
		log.Warn(fmt.Sprintf("Unknown environment, used: %s", envProd))
	}

	return log
}
