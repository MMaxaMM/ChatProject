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
	audioChatIdHeader = "chat_id"
	audioKey          = "audio"
)

type AudioHandler struct {
	service *service.Service
	log     *slog.Logger
}

func NewAudioHandler(service *service.Service, log *slog.Logger) *AudioHandler {
	return &AudioHandler{service: service, log: log}
}

func (h *AudioHandler) Recognize(c *gin.Context) {
	const op = "handler.Recognize"
	log := h.log.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		log.Error("Failed to resolve user id", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	chatIdStr := c.Request.URL.Query().Get(audioChatIdHeader)
	chatId, err := strconv.ParseInt(chatIdStr, 10, 64)
	if err != nil {
		log.Error("Bad request", slogx.Error(err))
		NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
		return
	}
	request := new(models.AudioRequest)
	request.UserId = userId
	request.ChatId = chatId

	src, hdr, err := c.Request.FormFile(audioKey)
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

	response, err := h.service.Recognize(request)
	if err != nil {
		if errors.Is(err, chat.ErrForeignKey) {
			log.Error("Chat not found", slogx.Error(err))
			NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
			return
		}
		if errors.Is(err, chat.ErrServiceNotAvailable) {
			log.Error("Audio service not available", slogx.Error(err))
			NewErrorResponse(c, http.StatusInternalServerError, MsgServiceNotAvailable)
			return
		}
		log.Error("Recognize error", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	log.Info("Recognition completed")

	c.JSON(http.StatusOK, response)
}
