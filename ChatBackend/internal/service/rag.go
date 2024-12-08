package service

import (
	"chat"
	ragv1 "chat/gen/rag"
	"chat/internal/config"
	"chat/internal/lib/markdown"
	"chat/internal/models"
	"chat/internal/repository"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RAGService struct {
	cfg config.RAG
	rep *repository.Repository
}

func NewRAGService(cfg config.RAG, rep *repository.Repository) *RAGService {
	return &RAGService{cfg: cfg, rep: rep}
}

func (s *RAGService) SendMessageRAG(request *models.RAGRequest) (*models.RAGResponse, error) {
	const op = "service.SendMessageRAG"

	request.Content = markdown.Prepare(request.Content)
	userId := request.UserId
	chatId := request.ChatId

	conn, err := grpc.NewClient(
		s.cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer conn.Close()

	client := ragv1.NewRAGServiceClient(conn)

	RAGResponse, err := client.Generate(
		context.Background(),
		&ragv1.RAGRequest{Content: request.Content},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", op, chat.ErrServiceNotAvailable, err)
	}

	response := &models.RAGResponse{
		UserId: userId,
		ChatId: chatId,
		Message: models.Message{
			Role:        models.RoleAssistant,
			Content:     markdown.Prepare(RAGResponse.Content),
			ContentType: models.TextType,
		},
	}

	// Сохранение запроса пользователя
	err = s.rep.SaveMessage(
		userId,
		chatId,
		models.RoleUser,
		request.Content,
		models.TextType,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Сохранение ответа сервиса
	err = s.rep.SaveMessage(
		userId,
		chatId,
		models.RoleAssistant,
		RAGResponse.Content,
		models.TextType,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response, nil
}
