package handler

import (
	"chat/internal/lib/slogx"
	"chat/internal/models"
	"chat/internal/service"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControlHandler struct {
	service *service.Service
	log     *slog.Logger
}

func NewControlHandler(service *service.Service, log *slog.Logger) *ControlHandler {
	return &ControlHandler{service: service, log: log}
}

func (h *ControlHandler) CreateChat(c *gin.Context) {
	const op = "handler.CreateChat"
	log := h.log.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		log.Error("Failed to resolve user id", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	request := new(models.CreateRequest)
	if err := c.BindJSON(request); err != nil {
		log.Error("Bad request", slogx.Error(err))
		NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
		return
	}
	request.UserId = userId

	log = log.With(
		slog.Int64("user_id", request.UserId),
		slog.Int("chat_type", int(request.ChatType)),
	)

	response, err := h.service.CreateChat(request)
	if err != nil {
		log.Error("Failed to create chat", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	log.Info("Chat was created", slog.Int64("chat_id", response.ChatId))

	c.JSON(http.StatusOK, response)
}

func (h *ControlHandler) DeleteChat(c *gin.Context) {
	const op = "handler.DeleteChat"
	log := h.log.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		log.Error("Failed to resolve user id", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	request := new(models.DeleteRequest)
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

	response, err := h.service.DeleteChat(request)
	if err != nil {
		log.Error("Failed to delete chat", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	log.Info("Chat was deleted")

	c.JSON(http.StatusOK, response)
}

func (h *ControlHandler) GetStart(c *gin.Context) {
	const op = "handler.GetStart"
	log := h.log.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		log.Error("Failed to resolve user id", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	request := new(models.StartRequest)
	request.UserId = userId

	log = log.With(slog.Int64("user_id", userId))

	response, err := h.service.GetStart(request)
	if err != nil {
		log.Error("Failed to get start page", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	log.Info("Get start page")

	c.JSON(http.StatusOK, response)
}

func (h *ControlHandler) GetHistory(c *gin.Context) {
	const op = "handler.GetHistory"
	log := h.log.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		log.Error("Failed to resolve user id", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	request := new(models.HistoryRequest)
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

	response, err := h.service.GetHistory(request)
	if err != nil {
		log.Error("Failed to get history", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	log.Info("History received")

	c.JSON(http.StatusOK, response)
}
