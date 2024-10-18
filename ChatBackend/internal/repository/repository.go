package repository

import (
	"chat"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(username, password string) (int, error)
	GetUser(username, password string) (*chat.User, error)
}

type ChatInterface interface {
	GetHistory(request *chat.HistoryRequest, limit int) (*chat.HistoryResponse, error)
	SaveChatItem(item *chat.ChatItem) error
	DeleteChat(request *chat.HistoryRequest) error
}

type Repository struct {
	Authorization
	ChatInterface
}

func NewPostgresRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ChatInterface: NewChatInterfacePostgres(db),
	}
}
