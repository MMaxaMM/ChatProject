package minioclient

import (
	"chat/internal/config"
	"chat/internal/models"
	"fmt"
	"time"

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

const expires = time.Second * 60 * 5

var Client *minio.Client

func Connect(cfg config.Minio) error {
	const op = "minioclient.Connect"

	var err error
	Client, err = minio.New(cfg.Address, cfg.User, cfg.Password, false)
	if err != err {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = MakeBucket(AudioBucketName)
	if err != err {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = MakeBucket(VideoBucketName)
	if err != err {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func MakeBucket(bucketName string) error {
	const op = "minioclient.MakeBucket"

	exists, err := Client.BucketExists(bucketName)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		err = Client.MakeBucket(bucketName, "us-east-1")
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}

func UploadObject(
	objectName string,
	object *models.Object,
	bucketName string,
	contentType ContentType,
) error {
	const op = "minioclient.UploadObject"

	_, err := Client.PutObject(
		bucketName,
		objectName,
		object.Payload,
		object.PayloadSize,
		minio.PutObjectOptions{ContentType: string(contentType)},
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func DeleteObject(
	bucketName string,
	objectName string,
) error {
	const op = "minioclient.DeleteObject"

	err := Client.RemoveObject(bucketName, objectName)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func GetURI(
	bucketName string,
	objectName string,
) (string, error) {
	const op = "minioclient.GetURI"

	URI, err := Client.PresignedGetObject(bucketName, objectName, expires, nil)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return URI.String(), nil
}
