package service

import (
	"chat"
	audiov1 "chat/gen/audio"
	"chat/internal/config"
	"chat/internal/lib/slogx"
	"chat/internal/logger"
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
	cfg config.Audio
	rep *repository.Repository
}

func NewAudioService(
	cfg config.Audio,
	rep *repository.Repository,
) *AudioService {
	return &AudioService{cfg: cfg, rep: rep}
}

func (s *AudioService) Recognize(request *models.AudioRequest) (*models.AudioResponse, error) {
	const op = "service.Recognize"

	objectName := uuid.New().String() + ".mp3"

	objectPath, err := minioclient.UploadObject(
		objectName,
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

	audioRequest := &audiov1.AudioRequest{Filepath: objectPath}
	audioResponse, err := client.Recognize(context.Background(), audioRequest)
	if err != nil {
		go minioclient.DeleteObject(minioclient.AudioBucketName, objectName)
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
			Content:     objectPath,
			ContentType: models.AudioType,
		},
	)
	if err != nil {
		go func() {
			err = minioclient.DeleteObject(minioclient.AudioBucketName, objectName)
			if err != nil {
				logger.Logger.Warn("Faild to delete file from storfge", slogx.Error(err))
			}
		}()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.rep.SaveMessage(response.UserId, response.ChatId, &response.Message)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response, nil
}
