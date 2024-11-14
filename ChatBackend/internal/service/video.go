package service

import "chat/internal/repository"

type VideoRepository interface {
}

type VideoService struct {
	rep *repository.Repository
}

func NewVideoService(rep *repository.Repository) *VideoService {
	return &VideoService{rep: rep}
}
