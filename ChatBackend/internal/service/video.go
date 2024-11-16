package service

import (
	"chat"
	videov1 "chat/gen/video"
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

type VideoService struct {
	cfg   config.Video
	rep   *repository.Repository
	minio *minioclient.MinioProvider
}

func NewVideoService(
	cfg config.Video,
	rep *repository.Repository,
	minio *minioclient.MinioProvider,
) *VideoService {
	return &VideoService{cfg: cfg, rep: rep, minio: minio}
}

func (s *VideoService) Detect(request *models.VideoRequest) (*models.VideoResponse, error) {
	const op = "service.Detect"

	filename := uuid.New().String() + ".mp4"

	uri, err := s.minio.UploadObject(
		filename,
		&request.Object,
		minioclient.VideoBucketName,
		minioclient.VideoContentType,
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

	client := videov1.NewVideoServiceClient(conn)

	videoRequest := &videov1.VideoRequest{Uri: uri}
	videoResponse, err := client.Detect(context.Background(), videoRequest)
	if err != nil {
		return nil, fmt.Errorf("%s: %w: %w", op, chat.ErrServiceNotAvailable, err)
	}

	response := &models.VideoResponse{
		UserId: request.UserId,
		ChatId: request.ChatId,
		Message: models.Message{
			Role:    models.RoleAssistant,
			Content: videoResponse.Uri,
		},
	}

	err = s.rep.SaveMessage(
		request.UserId,
		request.ChatId,
		&models.Message{
			Role:    models.RoleUser,
			Content: uri,
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
