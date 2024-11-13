package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type VideoService interface {
}

type VideoHandler struct {
	service VideoService
	log     *slog.Logger
}

func NewVideoHandler(service VideoService, log *slog.Logger) *VideoHandler {
	return &VideoHandler{service: service, log: log}
}

func (h *VideoHandler) GetHistory(c *gin.Context) {
	panic("implement me")
}

func (h *VideoHandler) Detect(c *gin.Context) {
	panic("implement me")
}
