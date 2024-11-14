package service

import (
	"chat/internal/config"
	"chat/internal/models"
	"chat/internal/repository"
)

type Auth interface {
	CreateUser(request *models.SignUpRequest) (*models.SignUpResponse, error)
	GenerateToken(request *models.SignInRequest) (*models.SignInResponse, error)
}

type Middleware interface {
	ParseToken(accessToken string) (int64, error)
}

type Control interface {
	CreateChat(request *models.CreateRequest) (*models.CreateResponse, error)
	DeleteChat(request *models.DeleteRequest) error
	GetStart(request *models.StartRequest) (*models.StartResponse, error)
	GetHistory(request *models.HistoryRequest) (*models.HistoryResponse, error)
}

type Chat interface {
	SendMessage(request *models.ChatMessageRequest) (*models.ChatMessageResponse, error)
}

type Audio interface {
}

type Video interface {
}

type Service struct {
	Auth
	Middleware
	Control
	Chat
	Audio
	Video
}

func NewService(
	rep *repository.Repository,
	cfg *config.Config,
) *Service {
	return &Service{
		Auth:       NewAuthService(rep),
		Middleware: NewMiddlewareService(),
		Control:    NewControlService(rep),
		Chat:       NewChatService(rep, cfg.LLM),
		Audio:      NewAudioService(rep),
		Video:      NewVideoService(rep),
	}
}
