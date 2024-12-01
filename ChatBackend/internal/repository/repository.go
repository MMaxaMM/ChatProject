package repository

import (
	"chat/internal/models"

	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(username, password string) (int64, error)
	GetUserId(username, password string) (int64, error)
}

type Control interface {
	GetHistory(
		userId int64,
		chatId int64,
		visibleOnly bool,
		limit int,
	) ([]models.Message, error)
	SaveMessage(
		userId int64,
		chatId int64,
		role string,
		content string,
		contentType models.ContentType,
	) error
	DeleteChat(userId int64, chatId int64) error
	CreateChat(userId int64, chatType models.ChatType) (int64, error)
	GetStart(userId int64) ([]models.ChatPreview, error)
}

type Repository struct {
	Auth
	Control
}

func NewPostgresRepository(db *sqlx.DB, filestorage string) *Repository {
	return &Repository{
		Auth:    NewAuthPostgres(db),
		Control: NewControlPostgres(db, filestorage),
	}
}
