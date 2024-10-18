package main

import (
	"chat/internal/config"
	"chat/internal/repository"
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

	_ = repository.NewPostgresRepository(db)

	// Тесты, надо завернуть в тест

	// x, err := rep.GetUser("admin", "admin")
	// log.Println(x)
	// if err != nil {
	// 	log.Println(err)
	// 	z, ok := err.(*pq.Error)
	// 	if ok {
	// 		log.Println(z.Code)
	// 	}
	// }

	// err = rep.SaveChatItem(&chat.ChatItem{
	// 	UserId: 1,
	// 	ChatId: 1,
	// 	Message: chat.Message{
	// 		Role:    chat.RoleUser,
	// 		Content: "3",
	// 	},
	// })
	// if err != nil {
	// 	log.Println(err)
	// 	z, ok := err.(*pq.Error)
	// 	if ok {
	// 		log.Println(z.Code)
	// 	}
	// }

	// x, err := rep.GetHistory(&chat.HistoryRequest{UserId: 1, ChatId: 1}, repository.NoLimit)
	// log.Println(x)
	// if err != nil {
	// 	log.Println(err)
	// 	z, ok := err.(*pq.Error)
	// 	if ok {
	// 		log.Println(z.Code)
	// 	}
	// }

	// err = rep.DeleteChat(&chat.HistoryRequest{UserId: 1, ChatId: 1})
	// if err != nil {
	// 	log.Println(err)
	// 	z, ok := err.(*pq.Error)
	// 	if ok {
	// 		log.Println(z.Code)
	// 	}
	// }
}
