package repository

import (
	"chat/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/jmoiron/sqlx"
)

type ControlPostgres struct {
	db *sqlx.DB
}

func NewControlPostgres(db *sqlx.DB) *ControlPostgres {
	return &ControlPostgres{db: db}
}

func (r *ControlPostgres) CreateChat(userId int64, chatType models.ChatType) (int64, error) {
	const op = "repository.CreateChat"

	var chatId int64
	query := fmt.Sprintf("INSERT INTO %s (user_id, chat_type, date) values ($1, $2, $3) RETURNING id", chatsTable)

	row := r.db.QueryRow(query, userId, chatType, time.Now())
	if err := row.Scan(&chatId); err != nil {
		return 0, fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}

	return chatId, nil
}

func (r *ControlPostgres) DeleteChat(userId int64, chatId int64) error {
	const op = "repository.DeleteChat"

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", chatsTable)

	_, err := r.db.Exec(query, chatId, userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}

	return nil
}

func (r *ControlPostgres) GetStart(userId int64) ([]models.ChatPreview, error) {
	const op = "repository.GetStart"

	userChats := new([]models.UserChats)

	query := fmt.Sprintf("SELECT id, chat_type, date FROM %s WHERE user_id=$1 ORDER BY date DESC", chatsTable)
	err := r.db.Select(userChats, query, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}

	chatPreviewSlice := make([]models.ChatPreview, len(*userChats))

	for idx, userChat := range *userChats {
		chatPreview := new(models.ChatPreview)
		chatPreview.ChatType = userChat.ChatType
		chatPreview.ChatId = userChat.ChatId

		query := fmt.Sprintf("SELECT content, date FROM %s WHERE user_id=$1 AND chat_id=$2 ORDER BY date DESC", messagesTable)
		err = r.db.Get(chatPreview, query, userId, userChat.ChatId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				chatPreview.Date = userChat.Date
			} else {
				return nil, fmt.Errorf("%s: %w", op, PostgresNewError(err))
			}
		}

		chatPreviewSlice[idx] = *chatPreview
	}

	sort.Sort(models.ByDate(chatPreviewSlice))

	return chatPreviewSlice, nil
}

func (r *ControlPostgres) GetHistory(
	userId int64,
	chatId int64,
	visibleOnly bool,
	limit int,
) ([]models.Message, error) {
	const op = "repository.GetHistory"

	messages := make([]models.Message, 0)

	var query string
	if visibleOnly {
		query = fmt.Sprintf("SELECT role, content FROM %s WHERE user_id=$1 AND chat_id=$2 AND role!='%s' ORDER BY date ASC", messagesTable, models.RoleSystem)
	} else {
		query = fmt.Sprintf("SELECT role, content FROM %s WHERE user_id=$1 AND chat_id=$2 ORDER BY date ASC", messagesTable)
	}
	if limit >= 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}
	err := r.db.Select(&messages, query, userId, chatId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}
	return messages, nil
}

func (r *ControlPostgres) SaveMessage(userId int64, chatId int64, message *models.Message) error {
	const op = "repository.SaveChatItem"

	query := fmt.Sprintf("INSERT INTO %s (user_id, chat_id, date, role, content) values ($1, $2, $3, $4, $5)", messagesTable)

	_, err := r.db.Exec(query, userId, chatId, time.Now(), message.Role, message.Content)
	if err != nil {
		return fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}

	return nil
}
