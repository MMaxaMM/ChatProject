package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/MMaxaMM/ChatProject/Auth/internal/app"
	"github.com/MMaxaMM/ChatProject/Auth/internal/config"
	"github.com/MMaxaMM/ChatProject/Auth/internal/logger"
	"github.com/joho/godotenv"
)

const ENV_PATH = "./.open.env"

func main() {
	// Загрузка .env файла, расположенного локально:
	if err := godotenv.Load(ENV_PATH); err != nil {
		panic("Error loading env variables: " + err.Error())
	}

	// Инициализация конфигураций:
	cfg := config.MustLoad()

	// Инициализация логгера:
	log := logger.SetupLogger(cfg.Env)
	log.Info("Starting application", slog.String("env", cfg.Env))

	// Инициализация приложения
	application := app.New(log, cfg.Port, cfg.StoragePath, cfg.TokenTTL)

	// Запуск gRPC-сервера
	go application.GRPCServer.MustRun()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	s := <-stop

	application.GRPCServer.Stop()
	log.Info("Application stopped", slog.String("signal", s.String()))
}
