package service

import (
	"chat"
	llmv1 "chat/gen/llm"
	"chat/internal/config"
	"chat/internal/models"
	"chat/internal/repository"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var DefaultHistoryLimit int = 10
var DefaultMaxTokens uint32 = 512

type ChatService struct {
	cfg config.LLM
	rep *repository.Repository
}

func NewChatService(cfg config.LLM, rep *repository.Repository) *ChatService {
	DefaultHistoryLimit = cfg.HistoryLimit
	DefaultMaxTokens = cfg.MaxTokens

	return &ChatService{cfg: cfg, rep: rep}
}

func (s *ChatService) SendMessage(request *models.ChatRequest) (*models.ChatResponse, error) {
	const op = "service.SendMessage"

	messages, err := s.rep.GetHistory(request.UserId, request.ChatId, false, DefaultHistoryLimit)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	request.Message.ContentType = models.TextType
	messages = append(messages, request.Message)

	conn, err := grpc.NewClient(
		s.cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer conn.Close()

	client := llmv1.NewLLMServiceClient(conn)

	llmMessages := make([]*llmv1.Message, len(messages))
	for idx, message := range messages {
		llmMessages[idx] = &llmv1.Message{Role: message.Role, Content: message.Content}
	}

	llmRequest := &llmv1.LLMRequest{Messages: llmMessages, MaxTokens: DefaultMaxTokens}
	llmResponse, err := client.Generate(context.Background(), llmRequest)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", op, chat.ErrServiceNotAvailable, err)
	}

	response := &models.ChatResponse{
		UserId: request.UserId,
		ChatId: request.ChatId,
		Message: models.Message{
			Role:        llmResponse.Message.Role,
			Content:     llmResponse.Message.Content,
			ContentType: models.TextType,
		},
	}

	err = s.rep.SaveMessage(request.UserId, request.ChatId, &request.Message)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.rep.SaveMessage(response.UserId, response.ChatId, &response.Message)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response, nil
}
