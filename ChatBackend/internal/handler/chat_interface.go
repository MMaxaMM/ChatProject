package handler

import (
	"chat"
	"chat/internal/lib/slogx"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHistory(c *gin.Context) {
	const op = "handler.GetHistory"
	logger := h.logger.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		logger.Error("failed to resolve user id", slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request := new(chat.HistoryRequest)
	request.UserId = userId
	if err := c.BindJSON(request); err != nil {
		logger.Error("bad request", slogx.Error(err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.ChatInterface.GetHistory(request)
	if err != nil {
		logger.Error("failed to get history", slog.Int("user_id", request.UserId), slog.Int("chat_id", request.ChatId), slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Info("history received", slog.Int("user_id", request.UserId), slog.Int("chat_id", request.ChatId))
	c.JSON(http.StatusOK, response)
}

func (h *Handler) SendMessage(c *gin.Context) {
	const op = "handler.SendMessage"
	logger := h.logger.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		logger.Error("failed to resolve user id", slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request := new(chat.ChatItem)
	request.UserId = userId
	if err := c.BindJSON(request); err != nil {
		logger.Error("bad request", slogx.Error(err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.ChatInterface.SendMessage(request)
	if err != nil {
		logger.Error("messaging error", slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Info("messages were exchanged", slog.Int("user_id", request.UserId), slog.Int("chat_id", request.ChatId))
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Delete(c *gin.Context) {
	const op = "handler.Delete"
	logger := h.logger.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		logger.Error("failed to resolve user id", slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request := new(chat.HistoryRequest)
	request.UserId = userId
	if err := c.BindJSON(request); err != nil {
		logger.Error("bad request", slogx.Error(err))
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.services.ChatInterface.DeleteChat(request)
	if err != nil {
		logger.Error("failed to delete chat", slog.Int("user_id", request.UserId), slog.Int("chat_id", request.ChatId), slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Info("the chat was deleted", slog.Int("user_id", request.UserId), slog.Int("chat_id", request.ChatId))
	c.Status(http.StatusOK)
}
