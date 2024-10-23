package repository

import (
	"chat"
	"database/sql"
	"errors"
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
	if err != nil {
		return nil, PostgresNewError(err)
	}
	return response, nil
}

func (r *ChatInterfacePostgres) SaveChatItem(item *chat.ChatItem) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, chat_id, date, role, content) values ($1, $2, $3, $4, $5)", messagesTable)

	_, err := r.db.Exec(query, item.UserId, item.ChatId, time.Now(), item.Role, item.Content)
	if err != nil {
		return PostgresNewError(err)
	}

	return nil
}

func (r *ChatInterfacePostgres) DeleteChat(request *chat.HistoryRequest) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", chatsTable)

	_, err := r.db.Exec(query, request.ChatId, request.UserId)
	if err != nil {
		PostgresNewError(err)
	}
	return nil
}

func (r *ChatInterfacePostgres) CreateChat(request *chat.HistoryRequest) (int, error) {
	var chatId int
	query := fmt.Sprintf("INSERT INTO %s (user_id) values ($1) RETURNING id", chatsTable)

	row := r.db.QueryRow(query, request.UserId)
	if err := row.Scan(&chatId); err != nil {
		return 0, PostgresNewError(err)
	}

	return chatId, nil
}

func (r *ChatInterfacePostgres) GetStart(userId int) (*chat.StartResponse, error) {
	chatsId := new([]int)

	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id=$1 ORDER BY id ASC", chatsTable)
	err := r.db.Select(chatsId, query, userId)
	if err != nil {
		return nil, PostgresNewError(err)
	}

	startResponse := new(chat.StartResponse)
	startResponse.UserId = userId

	for _, chatId := range *chatsId {
		chatPreview := new(chat.ChatPreview)
		chatPreview.ChatId = chatId

		query := fmt.Sprintf("SELECT content FROM %s WHERE user_id=$1 AND chat_id=$2 ORDER BY date ASC", messagesTable)
		err := r.db.Get(&chatPreview.Content, query, userId, chatId)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return nil, PostgresNewError(err)
			}
		}

		startResponse.Chats = append(startResponse.Chats, *chatPreview)
	}

	return startResponse, nil
}
