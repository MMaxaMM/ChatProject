package handler

import (
	"chat"
	"chat/internal/lib/slogx"
	"chat/internal/models"
	"chat/internal/service"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	videoChatIdHeader = "chat_id"
	videoKey          = "video"
)

type VideoHandler struct {
	service *service.Service
	log     *slog.Logger
}

func NewVideoHandler(service *service.Service, log *slog.Logger) *VideoHandler {
	return &VideoHandler{service: service, log: log}
}

func (h *VideoHandler) Detect(c *gin.Context) {
	const op = "handler.Detect"
	log := h.log.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		log.Error("Failed to resolve user id", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	chatIdStr := c.Request.URL.Query().Get(videoChatIdHeader)
	chatId, err := strconv.ParseInt(chatIdStr, 10, 64)
	if err != nil {
		log.Error("Bad request", slogx.Error(err))
		NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
		return
	}
	request := new(models.VideoRequest)
	request.UserId = userId
	request.ChatId = chatId

	src, hdr, err := c.Request.FormFile(videoKey)
	if err != nil {
		log.Error("Bad request", slogx.Error(err))
		NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
		return
	}
	request.Object = models.Object{Payload: src, PayloadSize: hdr.Size}
	defer src.Close()

	log = log.With(
		slog.Int64("user_id", request.UserId),
		slog.Int64("chat_id", request.ChatId),
	)

	response, err := h.service.Detect(request)
	if err != nil {
		if errors.Is(err, chat.ErrForeignKey) {
			log.Error("Chat not found", slogx.Error(err))
			NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
			return
		}
		if errors.Is(err, chat.ErrServiceNotAvailable) {
			log.Error("Video service not available", slogx.Error(err))
			NewErrorResponse(c, http.StatusServiceUnavailable, MsgServiceNotAvailable)
			return
		}
		log.Error("Detection error", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	log.Info("Detection completed")

	c.JSON(http.StatusOK, response)
}
