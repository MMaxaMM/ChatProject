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

type ByDate []ChatPreview

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Less(i, j int) bool { return a[i].Date.After(a[j].Date) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ChatPreview struct {
	ChatType ChatType  `json:"chat_type"`
	ChatId   int64     `json:"chat_id"`
	Content  string    `json:"content" db:"content"`
	Date     time.Time `json:"date" db:"date"`
}

type UserChats struct {
	ChatId   int64     `json:"chat_id" db:"id"`
	ChatType ChatType  `json:"chat_type" db:"chat_type"`
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
	UserId int64 `json:"user_id"`
	ChatId int64 `json:"chat_id"`
}

// History:

type HistoryRequest struct {
	UserId int64 `json:"user_id"`
	ChatId int64 `json:"chat_id"`
}

type HistoryResponse struct {
	UserId   int64     `json:"user_id"`
	ChatId   int64     `json:"chat_id"`
	Messages []Message `json:"messages"`
}
