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
	cfg config.Video
	rep *repository.Repository
}

func NewVideoService(
	cfg config.Video,
	rep *repository.Repository,
) *VideoService {
	return &VideoService{cfg: cfg, rep: rep}
}

func (s *VideoService) Detect(request *models.VideoRequest) (*models.VideoResponse, error) {
	const op = "service.Detect"

	userId := request.UserId
	chatId := request.ChatId

	objectName := uuid.New().String() + ".mp4"

	err := minioclient.UploadObject(
		objectName,
		&request.Object,
		minioclient.VideoBucketName,
		minioclient.VideoContentType,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	URI, err := minioclient.GetURI(minioclient.VideoBucketName, objectName)
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

	videoRequest := &videov1.VideoRequest{URI: URI}
	videoResponse, err := client.Detect(context.Background(), videoRequest)
	if err != nil {
		go minioclient.DeleteObject(minioclient.VideoBucketName, objectName)
		return nil, fmt.Errorf("%s: %w: %w", op, chat.ErrServiceNotAvailable, err)
	}

	URI, err = minioclient.GetURI(minioclient.VideoBucketName, videoResponse.ObjectName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := &models.VideoResponse{
		UserId: userId,
		ChatId: chatId,
		Message: models.Message{
			Role:        models.RoleAssistant,
			Content:     URI,
			ContentType: models.VideoType,
		},
	}

	// Сохранение запроса пользователя
	err = s.rep.SaveMessage(
		userId,
		chatId,
		models.RoleUser,
		objectName,
		models.VideoType,
	)
	if err != nil {
		go minioclient.DeleteObject(minioclient.VideoBucketName, objectName)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Сохранение ответа сервиса
	err = s.rep.SaveMessage(
		userId,
		chatId,
		models.RoleAssistant,
		videoResponse.ObjectName,
		models.VideoType,
	)
	if err != nil {
		go minioclient.DeleteObject(minioclient.VideoBucketName, videoResponse.ObjectName)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response, nil
}
