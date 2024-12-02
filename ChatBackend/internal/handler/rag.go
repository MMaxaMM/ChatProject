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

type RAGHandler struct {
	service *service.Service
	log     *slog.Logger
}

func NewRAGHandler(service *service.Service, log *slog.Logger) *RAGHandler {
	return &RAGHandler{service: service, log: log}
}

func (h *RAGHandler) SendMessageRAG(c *gin.Context) {
	const op = "handler.SendMessageRAG"
	log := h.log.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		log.Error("Failed to resolve user id", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	request := new(models.RAGRequest)
	if err := c.BindJSON(request); err != nil {
		log.Error("Bad request", slogx.Error(err))
		NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
		return
	}
	request.UserId = userId

	log = log.With(
		slog.Int64("user_id", request.UserId),
		slog.Int64("chat_id", request.ChatId),
	)

	response, err := h.service.SendMessageRAG(request)
	if err != nil {
		if errors.Is(err, chat.ErrForeignKey) {
			log.Error("Chat not found", slogx.Error(err))
			NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
			return
		}
		if errors.Is(err, chat.ErrServiceNotAvailable) {
			log.Error("RAG service not available", slogx.Error(err))
			NewErrorResponse(c, http.StatusInternalServerError, MsgServiceNotAvailable)
			return
		}
		log.Error("Messaging error", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	log.Info("Messages were exchanged")

	c.JSON(http.StatusOK, response)
}
