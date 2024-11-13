package service

import "chat/internal/models"

type ControlRepository interface {
	CreateChat(*models.CreateRequest) (*models.CreateResponse, error)
	DeleteChat(*models.DeleteRequest) error
	GetStart(*models.StartRequest) (*models.StartResponse, error)
}

type ControlService struct {
	rep ControlRepository
}

func NewControlService(rep ControlRepository) *ControlService {
	return &ControlService{rep: rep}
}

func (s *ControlService) CreateChat(request *models.CreateRequest) (*models.CreateResponse, error) {
	return s.rep.CreateChat(request)
}

func (s *ControlService) DeleteChat(request *models.DeleteRequest) error {
	return s.rep.DeleteChat(request)
}

func (s *ControlService) GetStart(request *models.StartRequest) (*models.StartResponse, error) {
	return s.rep.GetStart(request)
}
