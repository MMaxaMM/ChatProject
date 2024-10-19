package main

import (
	"chat/internal/api/llmapi"
	"chat/internal/config"
	"chat/internal/repository"
	"chat/internal/service"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Загрузка .env файла, расположенного локально
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	// Загрузка конфигураций
	cfg := config.MustLoad()

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
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	// Инициализация репозитория
	rep := repository.NewPostgresRepository(db)

	// Инициализация внутренних сервисов
	client := llmapi.NewClient(cfg.URL)
	services := service.NewService(rep, client)
	_ = services
}
