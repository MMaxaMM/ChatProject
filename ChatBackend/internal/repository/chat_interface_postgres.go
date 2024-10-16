package repository

import (
	"chat"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const NoLimit = -1

type ChatInterfacePostgres struct {
	db *sqlx.DB
}

func NewChatInterfacePostgres(db *sqlx.DB) *ChatInterfacePostgres {
	return &ChatInterfacePostgres{db: db}
}

func (r *ChatInterfacePostgres) GetHistory(request *chat.HistoryRequest, limit int) (*chat.HistoryResponse, error) {
	response := &chat.HistoryResponse{
		UserId: request.UserId,
		ChatId: request.ChatId,
	}

	query := fmt.Sprintf("SELECT role, content FROM %s WHERE user_id=$1 AND chat_id=$2 ORDER BY date ASC", messagesTable)
	if limit >= 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}
	err := r.db.Select(&response.Messages, query, request.UserId, request.ChatId)

	return response, err
}

func (r *ChatInterfacePostgres) SaveChatItem(item *chat.ChatItem) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, chat_id, date, role, content) values ($1, $2, $3, $4, $5)", messagesTable)

	_, err := r.db.Exec(query, item.UserId, item.ChatId, time.Now(), item.Role, item.Content)

	return err
}

func (r *ChatInterfacePostgres) DeleteChat(request *chat.HistoryRequest) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND chat_id=$2", messagesTable)

	_, err := r.db.Exec(query, request.UserId, request.ChatId)

	return err
}
