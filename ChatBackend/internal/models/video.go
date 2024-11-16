package models

type VideoRequest struct {
	UserId int64 `json:"user_id"`
	ChatId int64 `json:"chat_id"`
	Object
}

type VideoResponse struct {
	UserId  int64 `json:"user_id"`
	ChatId  int64 `json:"chat_id"`
	Message `json:"message"`
}
