package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type AudioService interface {
}

type AudioHandler struct {
	service AudioService
	log     *slog.Logger
}

func NewAudioHandler(service AudioService, log *slog.Logger) *AudioHandler {
	return &AudioHandler{service: service, log: log}
}

func (h *AudioHandler) GetHistory(c *gin.Context) {
	panic("implement me")
}

func (h *AudioHandler) Recognize(c *gin.Context) {
	panic("implement me")
}
