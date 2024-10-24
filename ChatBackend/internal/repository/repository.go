package repository

import (
	"chat"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(username, password string) (int, error)
	GetUserId(username, password string) (int, error)
}

type ChatInterface interface {
	GetHistory(request *chat.HistoryRequest, limit int) (*chat.HistoryResponse, error)
	SaveChatItem(item *chat.ChatItem) error
	DeleteChat(request *chat.HistoryRequest) error
	CreateChat(request *chat.HistoryRequest) (int, error)
	GetStart(userId int) (*chat.StartResponse, error)
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
