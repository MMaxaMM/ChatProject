package minioclient

import (
	"chat/internal/config"
	"chat/internal/models"
	"fmt"

	"github.com/minio/minio-go"
)

const (
	AudioBucketName = "audio"
	VideoBucketName = "video"
)

type ContentType string

const (
	AudioContentType ContentType = "audio/mpeg"
	VideoContentType ContentType = "video/mp4"
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

	exists, err := m.client.BucketExists(AudioBucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = m.client.MakeBucket(AudioBucketName, "us-east-1")
		if err != nil {
			return err
		}
	}

	exists, err = m.client.BucketExists(VideoBucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = m.client.MakeBucket(VideoBucketName, "us-east-1")
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MinioProvider) UploadObject(
	filename string,
	object *models.Object,
	bucketName string,
	contentType ContentType,
) (string, error) {
	const op = "minioclient.UploadAudio"

	_, err := m.client.PutObject(
		bucketName,
		filename,
		object.Payload,
		object.PayloadSize,
		minio.PutObjectOptions{ContentType: string(contentType)},
	)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return fmt.Sprintf("http://%s/%s/%s", m.cfg.Address, bucketName, filename), nil
}
