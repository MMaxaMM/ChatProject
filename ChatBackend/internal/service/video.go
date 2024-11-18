package service

import (
	"chat"
	videov1 "chat/gen/video"
	"chat/internal/config"
	"chat/internal/lib/slogx"
	"chat/internal/logger"
	minioclient "chat/internal/minio-client"
	"chat/internal/models"
	"chat/internal/repository"
	"context"
	"fmt"
	"strings"

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

	objectName := uuid.New().String() + ".mp4"

	objectPath, err := minioclient.UploadObject(
		objectName,
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

	videoRequest := &videov1.VideoRequest{Filepath: objectPath}
	videoResponse, err := client.Detect(context.Background(), videoRequest)
	if err != nil {
		go minioclient.DeleteObject(minioclient.VideoBucketName, objectName)
		return nil, fmt.Errorf("%s: %w: %w", op, chat.ErrServiceNotAvailable, err)
	}

	response := &models.VideoResponse{
		UserId: request.UserId,
		ChatId: request.ChatId,
		Message: models.Message{
			Role:        models.RoleAssistant,
			Content:     videoResponse.Filepath,
			ContentType: models.VideoType,
		},
	}

	err = s.rep.SaveMessage(
		request.UserId,
		request.ChatId,
		&models.Message{
			Role:        models.RoleUser,
			Content:     objectPath,
			ContentType: models.VideoType,
		},
	)
	if err != nil {
		go func() {
			err = minioclient.DeleteObject(minioclient.VideoBucketName, objectName)
			if err != nil {
				logger.Logger.Warn("Faild to delete file from storfge", slogx.Error(err))
			}
		}()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.rep.SaveMessage(response.UserId, response.ChatId, &response.Message)
	if err != nil {
		go func() {
			bucketNameAndObjectName := strings.Split(videoResponse.Filepath, "/")
			err = minioclient.DeleteObject(
				bucketNameAndObjectName[0],
				bucketNameAndObjectName[1],
			)
			if err != nil {
				logger.Logger.Warn("Faild to delete file from storfge", slogx.Error(err))
			}
		}()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return response, nil
}
