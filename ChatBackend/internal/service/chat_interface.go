package service

import (
	"chat"
	"chat/internal/api/llmapi"
	"chat/internal/repository"
)

var DefaultHistoryLimit int = 10
var DefaultMaxTokens uint = 512

type ChatInterfaceService struct {
	rep    repository.ChatInterface
	client *llmapi.Client
}

func NewChatInterfaceService(rep repository.ChatInterface, client *llmapi.Client) *ChatInterfaceService {
	return &ChatInterfaceService{rep: rep, client: client}
}

func (s *ChatInterfaceService) GetHistory(request *chat.HistoryRequest) (*chat.HistoryResponse, error) {
	return s.rep.GetHistory(request, repository.NoLimit)
}

func (s *ChatInterfaceService) DeleteChat(request *chat.HistoryRequest) error {
	return s.rep.DeleteChat(request)
}

func (s *ChatInterfaceService) SendMessage(item *chat.ChatItem) (*chat.ChatItem, error) {
	historyRequest := &chat.HistoryRequest{UserId: item.UserId, ChatId: item.ChatId}
	history, err := s.rep.GetHistory(historyRequest, DefaultHistoryLimit)
	if err != nil {
		return nil, err
	}

	messages := append(history.Messages, item.Message)
	request := llmapi.Request{Messages: messages, MaxTokens: DefaultMaxTokens}
	response, err := s.client.Generate(&request)
	if err != nil {
		return nil, err
	}

	err = s.rep.SaveChatItem(item)
	if err != nil {
		return nil, err
	}

	item.Message = *response
	err = s.rep.SaveChatItem(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *ChatInterfaceService) CreateChat(request *chat.HistoryRequest) (int, error) {
	return s.rep.CreateChat(request)
}

func (s *ChatInterfaceService) GetStart(userId int) (*chat.StartResponse, error) {
	return s.rep.GetStart(userId)
}
