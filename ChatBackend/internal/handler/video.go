package handler

import (
	"chat/internal/service"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	service *service.Service
	log     *slog.Logger
}

func NewVideoHandler(service *service.Service, log *slog.Logger) *VideoHandler {
	return &VideoHandler{service: service, log: log}
}

func (h *VideoHandler) GetHistory(c *gin.Context) {
	panic("implement me")
}

func (h *VideoHandler) Detect(c *gin.Context) {
	panic("implement me")
}
