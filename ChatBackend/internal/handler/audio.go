package handler

import (
	"chat"
	"chat/internal/lib/slogx"
	"chat/internal/models"
	"chat/internal/service"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

const audioKey = "audio"

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

	request := new(models.AudioRequest)
	if err := c.BindJSON(request); err != nil {
		log.Error("Bad request", slogx.Error(err))
		NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
		return
	}
	request.UserId = userId

	src, hdr, err := c.Request.FormFile(audioKey)
	if err != nil {
		log.Error("Bad request", slogx.Error(err))
		NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
		return
	}
	request.Audio = models.Audio{Payload: src, PayloadSize: hdr.Size}
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
