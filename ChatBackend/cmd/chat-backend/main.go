package main

import (
	"chat/internal/app"
	"chat/internal/config"
	"chat/internal/logger"
	"fmt"
	"os"
)

func main() {

	// Загрузка конфигураций
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading configuration file: %s\n", err.Error())
		os.Exit(1)
	}

	// Инициализация logger
	log := logger.SetupLogger(cfg.Env)
	log.Debug("Running in debug mode")

	// Запуск приложения
	application := app.NewApp(cfg, log)
	application.MustRun()
}
