package models

// Message:

type ChatMessageRequest struct {
	UserId  int64 `json:"user_id"`
	ChatId  int64 `json:"chat_id"`
	Message `json:"message"`
}

type ChatMessageResponse struct {
	UserId  int64 `json:"user_id"`
	ChatId  int64 `json:"chat_id"`
	Message `json:"message"`
}
