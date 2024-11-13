package repository

import (
	"chat/internal/models"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ControlPostgres struct {
	db *sqlx.DB
}

func NewControlPostgres(db *sqlx.DB) *ControlPostgres {
	return &ControlPostgres{db: db}
}

func (r *ControlPostgres) CreateChat(request *models.CreateRequest) (*models.CreateResponse, error) {
	const op = "repository.CreateChat"

	response := &models.CreateResponse{
		UserId:   request.UserId,
		ChatType: request.ChatType,
	}
	query := fmt.Sprintf("INSERT INTO %s (user_id, chat_type) values ($1, $2) RETURNING id", chatsTable)

	row := r.db.QueryRow(query, request.UserId, request.ChatType)
	if err := row.Scan(&response.ChatId); err != nil {
		return nil, fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}

	return response, nil
}

func (r *ControlPostgres) DeleteChat(request *models.DeleteRequest) error {
	const op = "repository.DeleteChat"

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2 AND chat_type=$3", chatsTable)

	_, err := r.db.Exec(query, request.ChatId, request.UserId, request.ChatType)
	if err != nil {
		return fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}

	return nil
}

func (r *ControlPostgres) GetStart(request *models.StartRequest) (*models.StartResponse, error) {
	const op = "repository.GetStart"

	userChats := new([]models.UserChats)

	query := fmt.Sprintf("SELECT id, chat_type FROM %s WHERE user_id=$1 ORDER BY id ASC", chatsTable)
	err := r.db.Select(userChats, query, request.UserId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, PostgresNewError(err))
	}

	response := new(models.StartResponse)
	response.UserId = request.UserId

	for _, userChat := range *userChats {
		chatPreview := new(models.ChatPreview)
		chatPreview.ChatType = userChat.ChatType
		chatPreview.ChatId = userChat.ChatId

		switch chatPreview.ChatType {
		case models.TypeChat:
			query := fmt.Sprintf("SELECT content, date FROM %s WHERE user_id=$1 AND chat_id=$2 ORDER BY date DESC", chatTable)
			err = r.db.Get(chatPreview, query, request.UserId, userChat.ChatId)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return nil, fmt.Errorf("%s: %w", op, PostgresNewError(err))
				}
			}
		case models.TypeRAG:
			query := fmt.Sprintf("SELECT content, date FROM %s WHERE user_id=$1 AND chat_id=$2 ORDER BY date DESC", ragTable)
			err = r.db.Get(chatPreview, query, request.UserId, userChat.ChatId)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return nil, fmt.Errorf("%s: %w", op, PostgresNewError(err))
				}
			}
		case models.TypeAudio:
			query := fmt.Sprintf("SELECT date FROM %s WHERE user_id=$1 AND chat_id=$2 ORDER BY date DESC", audioTable)
			err = r.db.Get(chatPreview, query, request.UserId, userChat.ChatId)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return nil, fmt.Errorf("%s: %w", op, PostgresNewError(err))
				} else {
					chatPreview.Content = audioPlug
				}
			}
		case models.TypeVideo:
			query := fmt.Sprintf("SELECT date FROM %s WHERE user_id=$1 AND chat_id=$2 ORDER BY date DESC", videoTable)
			err = r.db.Get(chatPreview, query, request.UserId, userChat.ChatId)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return nil, fmt.Errorf("%s: %w", op, PostgresNewError(err))
				} else {
					chatPreview.Content = videoPlug
				}
			}
		default:
			return nil, fmt.Errorf("%s: %w", op, errors.New("unknown chat type"))
		}
		response.Chats = append(response.Chats, *chatPreview)
	}

	return response, nil
}
