package service

import (
	"chat"
	audiov1 "chat/gen/audio"
	"chat/internal/config"
	"chat/internal/lib/markdown"
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

	userId := request.UserId
	chatId := request.ChatId

	objectName := uuid.New().String() + ".mp3"

	err := minioclient.UploadObject(
		objectName,
		&request.Object,
		minioclient.AudioBucketName,
		minioclient.AudioContentType,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	URI, err := minioclient.GetURI(minioclient.AudioBucketName, objectName)
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

	audioRequest := &audiov1.AudioRequest{URI: URI}
	audioResponse, err := client.Recognize(context.Background(), audioRequest)
	if err != nil {
		go minioclient.DeleteObject(minioclient.AudioBucketName, objectName)
		return nil, fmt.Errorf("%s: %w: %w", op, chat.ErrServiceNotAvailable, err)
	}
	audioResponse.Content = markdown.Prepare(audioResponse.Content)

	response := &models.AudioResponse{
		UserId: userId,
		ChatId: chatId,
		Message: models.Message{
			Role:        models.RoleAssistant,
			Content:     audioResponse.Content,
			ContentType: models.TextType,
		},
	}

	// Сохранение запроса пользователя
	err = s.rep.SaveMessage(
		userId,
		chatId,
		models.RoleUser,
		objectName,
		models.AudioType,
	)
	if err != nil {
		minioclient.DeleteObject(minioclient.AudioBucketName, objectName)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Сохранение ответа сервиса
	err = s.rep.SaveMessage(
		userId,
		chatId,
		models.RoleAssistant,
		response.Content,
		models.TextType,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response, nil
}
