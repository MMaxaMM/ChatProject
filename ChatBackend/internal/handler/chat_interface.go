package handler

import (
	"chat"
	"chat/internal/lib/slogx"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHistory(c *gin.Context) {
	//const op = "handler.GetHistory"
	logger := h.logger //.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		logger.Error("failed to resolve user id", slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request := new(chat.HistoryRequest)
	if err := c.BindJSON(request); err != nil {
		logger.Error("bad request", slogx.Error(err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	request.UserId = userId

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
	//const op = "handler.SendMessage"
	logger := h.logger //.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		logger.Error("failed to resolve user id", slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request := new(chat.ChatItem)
	if err := c.BindJSON(request); err != nil {
		logger.Error("bad request", slogx.Error(err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	request.UserId = userId

	response, err := h.services.ChatInterface.SendMessage(request)
	if err != nil {
		switch chat.ErrorCode(err) {
		case chat.EFOREIGNKEY:
			logger.Error("incorrect chat_id", slog.Int("chat_id", request.ChatId))
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		default:
			logger.Error("messaging error", slogx.Error(err))
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	logger.Info("messages were exchanged", slog.Int("user_id", request.UserId), slog.Int("chat_id", request.ChatId))
	c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteChat(c *gin.Context) {
	//const op = "handler.DeleteChat"
	logger := h.logger //.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		logger.Error("failed to resolve user id", slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request := new(chat.HistoryRequest)
	if err := c.BindJSON(request); err != nil {
		logger.Error("bad request", slogx.Error(err))
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	request.UserId = userId

	err = h.services.ChatInterface.DeleteChat(request)
	if err != nil {
		logger.Error("failed to delete chat", slog.Int("user_id", request.UserId), slog.Int("chat_id", request.ChatId), slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Info("the chat was deleted", slog.Int("user_id", request.UserId), slog.Int("chat_id", request.ChatId))
	c.Status(http.StatusOK)
}

func (h *Handler) CreateChat(c *gin.Context) {
	//const op = "handler.CreateChat"
	logger := h.logger //.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		logger.Error("failed to resolve user id", slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request := new(chat.HistoryRequest)
	request.UserId = userId

	chatId, err := h.services.ChatInterface.CreateChat(request)
	if err != nil {
		logger.Error("failed to create chat", slog.Int("user_id", userId), slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	request.ChatId = chatId

	logger.Info("the chat was created", slog.Int("user_id", request.UserId), slog.Int("chat_id", request.ChatId))
	c.JSON(http.StatusOK, request)
}

func (h *Handler) GetStart(c *gin.Context) {
	//const op = "handler.GetStart"
	logger := h.logger //.With(slog.String("op", op))

	userId, err := getUserId(c)
	if err != nil {
		logger.Error("failed to resolve user id", slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := h.services.GetStart(userId)
	if err != nil {
		logger.Error("failed to get start page", slog.Int("user_id", userId), slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Info("get start page", slog.Int("user_id", userId))
	c.JSON(http.StatusOK, response)
}
