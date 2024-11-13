package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "public.users"
	chatsTable = "public.chats"
	chatTable  = "public.chat"
	ragTable   = "public.rag"
	audioTable = "public.audio"
	videoTable = "public.video"
)

const (
	audioPlug = "Audio"
	videoPlug = "Video"
)

const NoLimit = -1

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
