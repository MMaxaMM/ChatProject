package repository

import (
	"chat/internal/models"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type ChatPostgres struct {
	db *sqlx.DB
}

func NewChatPostgres(db *sqlx.DB) *ChatPostgres {
	return &ChatPostgres{db: db}
}

func (r *ChatPostgres) GetHistory(userId int64, chatId int64, limit int) ([]models.Message, error) {
	const op = "repository.GetHistory"

	messages := make([]models.Message, 0)

	query := fmt.Sprintf("SELECT role, content FROM %s WHERE user_id=$1 AND chat_id=$2 ORDER BY date ASC", chatTable)
	if limit >= 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}
	err := r.db.Select(&messages, query, userId, chatId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}
	return messages, nil
}

func (r *ChatPostgres) SaveChatItem(userId int64, chatId int64, message models.Message) error {
	const op = "repository.SaveChatItem"

	query := fmt.Sprintf("INSERT INTO %s (user_id, chat_id, date, role, content) values ($1, $2, $3, $4, $5)", chatTable)

	_, err := r.db.Exec(query, userId, chatId, time.Now(), message.Role, message.Content)
	if err != nil {
		return fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}

	return nil
}
