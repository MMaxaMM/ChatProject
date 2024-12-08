package service

import (
	"chat/internal/models"
	"chat/internal/repository"
	"fmt"
)

type ControlService struct {
	rep *repository.Repository
}

func NewControlService(rep *repository.Repository) *ControlService {
	return &ControlService{rep: rep}
}

func (s *ControlService) CreateChat(request *models.CreateRequest) (*models.CreateResponse, error) {
	const op = "service.CreateChat"

	chatId, err := s.rep.CreateChat(request.UserId, request.ChatType)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := &models.CreateResponse{
		UserId:   request.UserId,
		ChatType: request.ChatType,
		ChatId:   chatId,
	}

	return response, nil
}

func (s *ControlService) DeleteChat(request *models.DeleteRequest) (*models.DeleteResponse, error) {
	const op = "service.DeleteChat"

	err := s.rep.DeleteChat(request.UserId, request.ChatId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := &models.DeleteResponse{UserId: request.UserId, ChatId: request.ChatId}

	return response, nil
}

func (s *ControlService) GetStart(request *models.StartRequest) (*models.StartResponse, error) {
	const op = "service.GetStart"

	chatPreviewSlice, err := s.rep.GetStart(request.UserId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := &models.StartResponse{UserId: request.UserId, Chats: chatPreviewSlice}

	return response, nil
}

func (s *ControlService) GetHistory(request *models.HistoryRequest) (*models.HistoryResponse, error) {
	const op = "service.GetHistory"

	messages, err := s.rep.GetHistory(
		request.UserId,
		request.ChatId,
		true,
		repository.NoLimit,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := &models.HistoryResponse{
		UserId:   request.UserId,
		ChatId:   request.ChatId,
		Messages: messages,
	}

	return response, nil
}
