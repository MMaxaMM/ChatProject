package models

type AudioRequest struct {
	UserId int64 `json:"user_id"`
	ChatId int64 `json:"chat_id"`
	Object
}

type AudioResponse struct {
	UserId  int64 `json:"user_id"`
	ChatId  int64 `json:"chat_id"`
	Message `json:"message"`
}
