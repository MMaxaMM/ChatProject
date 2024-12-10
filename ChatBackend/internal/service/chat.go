package service

import (
	"chat"
	llmv1 "chat/gen/llm"
	"chat/internal/config"
	"chat/internal/lib/markdown"
	"chat/internal/models"
	"chat/internal/repository"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var HistoryLimit int
var MaxTokens uint32

type ChatService struct {
	cfg config.LLM
	rep *repository.Repository
}

func NewChatService(cfg config.LLM, rep *repository.Repository) *ChatService {
	HistoryLimit = cfg.HistoryLimit
	MaxTokens = cfg.MaxTokens

	return &ChatService{cfg: cfg, rep: rep}
}

func (s *ChatService) SendMessage(request *models.ChatRequest) (*models.ChatResponse, error) {
	const op = "service.SendMessage"

	request.Content = markdown.Prepare(request.Content)
	userId := request.UserId
	chatId := request.ChatId

	messages, err := s.rep.GetHistory(userId, chatId, false, HistoryLimit)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	request.Role = models.RoleUser
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

	llmRequest := &llmv1.LLMRequest{Messages: llmMessages, MaxTokens: MaxTokens}
	llmResponse, err := client.Generate(context.Background(), llmRequest)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", op, chat.ErrServiceNotAvailable, err)
	}

	response := &models.ChatResponse{
		UserId: userId,
		ChatId: chatId,
		Message: models.Message{
			Role:        models.RoleAssistant,
			Content:     markdown.Prepare(llmResponse.Content),
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
		llmResponse.Content,
		models.TextType,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response, nil
}
