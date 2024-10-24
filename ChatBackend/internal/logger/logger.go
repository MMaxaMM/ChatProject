package logger

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
)

const (
	envDevLinux   = "dev-linux"
	envDevWindows = "dev-windows"
	envProd       = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDevLinux:
		opts := &tint.Options{Level: slog.LevelDebug, TimeFormat: time.Kitchen}
		log = slog.New(tint.NewHandler(os.Stdout, opts))

	case envDevWindows:
		opts := &tint.Options{Level: slog.LevelDebug, TimeFormat: time.Kitchen}
		log = slog.New(tint.NewHandler(colorable.NewColorable(os.Stdout), opts))

	case envProd:
		opts := &slog.HandlerOptions{Level: slog.LevelInfo}
		log = slog.New(slog.NewJSONHandler(os.Stderr, opts))

	default:
		opts := &slog.HandlerOptions{Level: slog.LevelInfo}
		log = slog.New(slog.NewJSONHandler(os.Stderr, opts))

		log.Warn(fmt.Sprintf("unknown environment, used: %s", envProd))
	}

	return log
}
