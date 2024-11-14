package minioclient

import (
	"chat/internal/config"

	"github.com/minio/minio-go"
)

const (
	audioBucketName = "audio"
	videoBucketName = "video"
)

type Minio struct {
	client *minio.Client
	cfg    config.Minio
}

func NewMinio(cfg config.Minio) *Minio {
	return &Minio{cfg: cfg}
}

func (m *Minio) Connect() error {
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
