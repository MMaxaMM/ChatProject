package main

import (
	"chat"
	"chat/internal/api/llmapi"
	"chat/internal/config"
	"chat/internal/handler"
	"chat/internal/lib/slogx"
	"chat/internal/logger"
	"chat/internal/repository"
	"chat/internal/service"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Загрузка .env файла, расположенного локально
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	// Загрузка конфигураций
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error loading configuration file: %s", err.Error())
	}

	// Инициализация logger
	logger := logger.SetupLogger(cfg.Env)
	logger.Debug("running in debug mode")

	// Инициализация базы данных
	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Username: cfg.Username,
		DBName:   cfg.DBName,
		SSLMode:  cfg.SSLMode,
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logger.Error("failed to initialize db", slogx.Error(err))
		os.Exit(1)
	}

	// Инициализация репозитория
	rep := repository.NewPostgresRepository(db)

	// Инициализация сервисов
	client := llmapi.NewClient(cfg.URL)
	services := service.NewService(rep, client)

	// Переопределение значений по умолчанию
	service.DefaultHistoryLimit = cfg.HistoryLimit
	service.DefaultMaxTokens = cfg.MaxTokens

	// Инициализация обработчиков
	handlers := handler.NewHandler(services, logger)
	router := handlers.InitRoutes()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Запуск сервера
	srv := new(chat.Server)
	logger.Info("run HTTP server")
	if err = srv.Run(cfg.Address, router); err != nil {
		logger.Error("failed to run HTTP server", slogx.Error(err))
		os.Exit(1)
	}
}
