package service

type VideoRepository interface {
}

type VideoService struct {
	rep VideoRepository
}

func NewVideoService(rep VideoRepository) *VideoService {
	return &VideoService{rep: rep}
}
