package service

type AudioRepository interface {
}

type AudioService struct {
	rep AudioRepository
}

func NewAudioService(rep AudioRepository) *AudioService {
	return &AudioService{rep: rep}
}
