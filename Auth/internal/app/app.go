package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/MMaxaMM/ChatProject/Auth/internal/app/grpc"
	"github.com/MMaxaMM/ChatProject/Auth/internal/services"
	"github.com/MMaxaMM/ChatProject/Auth/internal/storage/sqlite"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	port int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// Инициализация слоя работы с данными
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	// Инициализация сервисного слоя
	authService := services.New(log, storage, tokenTTL)

	// Инициализания и регистрация gRPC обработчиков
	grpcApp := grpcapp.New(log, port, authService)

	return &App{GRPCServer: grpcApp}
}
