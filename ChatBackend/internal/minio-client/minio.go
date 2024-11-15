package minioclient

import (
	"chat/internal/config"
	"chat/internal/models"
	"fmt"

	"github.com/google/uuid"
	"github.com/minio/minio-go"
)

const (
	audioBucketName = "audio"
	videoBucketName = "video"
)

type MinioProvider struct {
	client *minio.Client
	cfg    config.Minio
}

func NewMinioProvider(cfg config.Minio) *MinioProvider {
	return &MinioProvider{cfg: cfg}
}

func (m *MinioProvider) Connect() error {
	var err error
	m.client, err = minio.New(m.cfg.Address, m.cfg.User, m.cfg.Password, false)
	if err != err {
		return err
	}

	exists, err := m.client.BucketExists(audioBucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = m.client.MakeBucket(audioBucketName, "us-east-1")
		if err != nil {
			return err
		}
	}

	exists, err = m.client.BucketExists(videoBucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = m.client.MakeBucket(videoBucketName, "us-east-1")
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MinioProvider) UploadAudio(audio *models.Audio) (string, error) {
	const op = "minioclient.UploadAudio"

	filename := uuid.New().String() + ".mp3"

	_, err := m.client.PutObject(
		audioBucketName,
		filename,
		audio.Payload,
		audio.PayloadSize,
		minio.PutObjectOptions{ContentType: "audio/mpeg"},
	)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return fmt.Sprintf("http://%s/%s/%s", m.cfg.Address, audioBucketName, filename), nil
}

func (m *MinioProvider) DownloadAudio(filename string) (*models.Audio, error) {
	const op = "minioclient.DownloadAudio"

	payload, err := m.client.GetObject(
		audioBucketName,
		filename,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer payload.Close()

	return &models.Audio{Payload: payload}, nil
}
