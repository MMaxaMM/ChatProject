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
	SendMessage(request *models.ChatRequest) (*models.ChatResponse, error)
}

type RAG interface {
	SendMessageRAG(request *models.RAGRequest) (*models.RAGResponse, error)
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
	RAG
	Audio
	Video
}

func NewService(
	cfg *config.Config,
	rep *repository.Repository,
) *Service {
	return &Service{
		Auth:       NewAuthService(rep),
		Middleware: NewMiddlewareService(),
		Control:    NewControlService(rep),
		Chat:       NewChatService(cfg.LLM, rep),
		RAG:        NewRAGService(cfg.RAG, rep),
		Audio:      NewAudioService(cfg.Audio, rep),
		Video:      NewVideoService(cfg.Video, rep),
	}
}
