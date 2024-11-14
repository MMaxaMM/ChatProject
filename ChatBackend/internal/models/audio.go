package models

import "io"

type AudioRequest struct {
	UserId int64 `json:"user_id"`
	ChatId int64 `json:"chat_id"`
	Audio
}

type AudioResponse struct {
	UserId  int64 `json:"user_id"`
	ChatId  int64 `json:"chat_id"`
	Message `json:"message"`
}

type Audio struct {
	Payload     io.Reader
	PayloadSize int64
}
