package handler

import (
	"chat/internal/service"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type AudioHandler struct {
	service *service.Service
	log     *slog.Logger
}

func NewAudioHandler(service *service.Service, log *slog.Logger) *AudioHandler {
	return &AudioHandler{service: service, log: log}
}

func (h *AudioHandler) GetHistory(c *gin.Context) {
	panic("implement me")
}

func (h *AudioHandler) Recognize(c *gin.Context) {
	panic("implement me")
}
