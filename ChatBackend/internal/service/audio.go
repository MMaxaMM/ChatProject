package service

import (
	"chat"
	audiov1 "chat/gen/audio"
	"chat/internal/config"
	minioclient "chat/internal/minio-client"
	"chat/internal/models"
	"chat/internal/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AudioService struct {
	cfg   config.Audio
	rep   *repository.Repository
	minio *minioclient.MinioProvider
}

func NewAudioService(
	cfg config.Audio,
	rep *repository.Repository,
	minio *minioclient.MinioProvider,
) *AudioService {
	return &AudioService{cfg: cfg, rep: rep, minio: minio}
}

func (s *AudioService) Recognize(request *models.AudioRequest) (*models.AudioResponse, error) {
	const op = "service.Recognize"

	filename := uuid.New().String() + ".mp3"

	filepath, err := s.minio.UploadObject(
		filename,
		&request.Object,
		minioclient.AudioBucketName,
		minioclient.AudioContentType,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	conn, err := grpc.NewClient(
		s.cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer conn.Close()

	client := audiov1.NewAudioServiceClient(conn)

	audioRequest := &audiov1.AudioRequest{Filepath: filepath}
	audioResponse, err := client.Recognize(context.Background(), audioRequest)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", op, chat.ErrServiceNotAvailable, err)
	}

	response := &models.AudioResponse{
		UserId: request.UserId,
		ChatId: request.ChatId,
		Message: models.Message{
			Role:        models.RoleAssistant,
			Content:     audioResponse.Result,
			ContentType: models.TextType,
		},
	}

	err = s.rep.SaveMessage(
		request.UserId,
		request.ChatId,
		&models.Message{
			Role:        models.RoleUser,
			Content:     filepath,
			ContentType: models.AudioType,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.rep.SaveMessage(response.UserId, response.ChatId, &response.Message)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response, nil
}
