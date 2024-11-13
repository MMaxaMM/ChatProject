package main

import (
	"chat/internal/app"
	"chat/internal/config"
	"chat/internal/logger"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Загрузка .env файла, расположенного локально
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading environment variables: %s\n", err.Error())
		os.Exit(1)
	}

	// Загрузка конфигураций
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading configuration file: %s\n", err.Error())
		os.Exit(1)
	}

	// Инициализация logger
	log := logger.SetupLogger(cfg.Env)
	log.Debug("Running in debug mode")

	application := app.NewApp(cfg, log)
	application.MustRun()
}
