package service

import "chat/internal/repository"

type AudioService struct {
	rep *repository.Repository
}

func NewAudioService(rep *repository.Repository) *AudioService {
	return &AudioService{rep: rep}
}
