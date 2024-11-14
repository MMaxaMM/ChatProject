package models

type ChatRequest struct {
	UserId  int64 `json:"user_id"`
	ChatId  int64 `json:"chat_id"`
	Message `json:"message"`
}

type ChatResponse struct {
	UserId  int64 `json:"user_id"`
	ChatId  int64 `json:"chat_id"`
	Message `json:"message"`
}
