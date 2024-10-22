package repository

import (
	"chat"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const (
	usersTable    = "public.users"
	messagesTable = "public.messages"
)

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

func PostgresNewError(err error) error {
	if err, ok := err.(*pq.Error); ok {
		switch err.Code {
		case "23505":
			return chat.NewError(chat.EDUPLICATE, err)
		}
	}
	return chat.NewError(chat.EINTERNAL, err)
}
