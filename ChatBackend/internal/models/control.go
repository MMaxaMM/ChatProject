package models

import "time"

type ChatType int

const (
	TypeChat  ChatType = 1
	TypeRAG   ChatType = 2
	TypeAudio ChatType = 3
	TypeVideo ChatType = 4
)

// Start:

type StartRequest struct {
	UserId int64 `json:"user_id"`
}

type StartResponse struct {
	UserId int64         `json:"user_id"`
	Chats  []ChatPreview `json:"chats"`
}

type UserChats struct {
	ChatId   int64    `json:"chat_id" db:"id"`
	ChatType ChatType `json:"chat_type" db:"chat_type"`
}

type ChatPreview struct {
	ChatType ChatType  `json:"chat_type"`
	ChatId   int64     `json:"chat_id"`
	Content  string    `json:"content" db:"content"`
	Date     time.Time `json:"date" db:"date"`
}

// Create:

type CreateRequest struct {
	UserId   int64    `json:"user_id"`
	ChatType ChatType `json:"chat_type"`
}

type CreateResponse struct {
	UserId   int64    `json:"user_id"`
	ChatType ChatType `json:"chat_type"`
	ChatId   int64    `json:"chat_id" db:"id"`
}

// Delete:

type DeleteRequest struct {
	UserId   int64    `json:"user_id"`
	ChatType ChatType `json:"chat_type"`
	ChatId   int64    `json:"chat_id"`
}
