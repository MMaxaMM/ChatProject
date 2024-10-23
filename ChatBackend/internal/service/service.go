package service

import (
	"chat"
	"chat/internal/api/llmapi"
	"chat/internal/repository"
)

type Authorization interface {
	CreateUser(username, password string) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ChatInterface interface {
	GetHistory(request *chat.HistoryRequest) (*chat.HistoryResponse, error)
	DeleteChat(request *chat.HistoryRequest) error
	SendMessage(item *chat.ChatItem) (*chat.ChatItem, error)
	CreateChat(request *chat.HistoryRequest) (int, error)
	GetStart(userId int) (*chat.StartResponse, error)
}

type Service struct {
	Authorization
	ChatInterface
}

func NewService(rep *repository.Repository, client *llmapi.Client) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
		ChatInterface: NewChatInterfaceService(rep.ChatInterface, client),
	}
}
