package main

import (
	"chat/internal/config"
	minioclient "chat/internal/minio-client"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

	if err = minioclient.Connect(cfg.Minio); err != nil {
		log.Fatalln(err.Error())
	}
}
