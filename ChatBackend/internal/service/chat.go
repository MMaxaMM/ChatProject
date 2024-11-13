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

type ChatRepository interface {
	GetHistory(userId int64, chatId int64, limit int) ([]models.Message, error)
	SaveChatItem(userId int64, chatId int64, message models.Message) error
}

type ChatService struct {
	rep ChatRepository
	cfg config.LLM
}

func NewChatService(rep ChatRepository, cfg config.LLM) *ChatService {
	DefaultHistoryLimit = cfg.HistoryLimit
	DefaultMaxTokens = cfg.MaxTokens

	return &ChatService{rep: rep, cfg: cfg}
}

func (s *ChatService) GetHistory(request *models.ChatHistoryRequest) (*models.ChatHistoryResponse, error) {
	const op = "service.GetHistory"

	messages, err := s.rep.GetHistory(request.UserId, request.ChatId, repository.NoLimit)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := &models.ChatHistoryResponse{
		UserId:   request.UserId,
		ChatId:   request.ChatId,
		Messages: messages,
	}

	return response, nil
}

func (s *ChatService) SendMessage(request *models.ChatMessageRequest) (*models.ChatMessageResponse, error) {
	const op = "service.SendMessage"

	messages, err := s.rep.GetHistory(request.UserId, request.ChatId, DefaultHistoryLimit)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	messages = append(messages, request.Message)

	conn, err := grpc.NewClient(
		s.cfg.ChatAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, chat.ErrServiceNotAvailable)
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
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := new(models.ChatMessageResponse)
	response.UserId = request.UserId
	response.ChatId = request.ChatId
	response.Message = models.Message{
		Role:    llmResponse.Message.Role,
		Content: llmResponse.Message.Content,
	}

	err = s.rep.SaveChatItem(request.UserId, request.ChatId, request.Message)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.rep.SaveChatItem(request.UserId, request.ChatId, response.Message)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response, nil
}
