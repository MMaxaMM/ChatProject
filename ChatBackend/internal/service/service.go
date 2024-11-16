package service

import (
	"chat/internal/config"
	minioclient "chat/internal/minio-client"
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
	SendMessage(request *models.ChatRequest) (*models.ChatResponse, error)
}

type Audio interface {
	Recognize(request *models.AudioRequest) (*models.AudioResponse, error)
}

type Video interface {
	Detect(request *models.VideoRequest) (*models.VideoResponse, error)
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
	cfg *config.Config,
	rep *repository.Repository,
	minio *minioclient.MinioProvider,
) *Service {
	return &Service{
		Auth:       NewAuthService(rep),
		Middleware: NewMiddlewareService(),
		Control:    NewControlService(rep),
		Chat:       NewChatService(cfg.LLM, rep),
		Audio:      NewAudioService(cfg.Audio, rep, minio),
		Video:      NewVideoService(cfg.Video, rep, minio),
	}
}
