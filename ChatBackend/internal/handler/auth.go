package handler

import (
	"chat"
	"chat/internal/lib/slogx"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	const op = "handler.signUp"
	logger := h.logger.With(slog.String("op", op))

	var user chat.User

	if err := c.BindJSON(&user); err != nil {
		logger.Error("bad request", slogx.Error(err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(user.Username, user.Password)
	if err != nil {
		switch chat.ErrorCode(err) {
		case chat.EDUPLICATE:
			logger.Warn("user already exists", slog.String("username", user.Username))
			newErrorResponse(c, http.StatusConflict, err.Error())
			return
		default:
			logger.Error("failed to create user", slogx.Error(err))
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	logger.Info("new user created", slog.String("username", user.Username))
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) signIn(c *gin.Context) {
	const op = "handler.signIn"
	logger := h.logger.With(slog.String("op", op))

	var user chat.User

	if err := c.BindJSON(&user); err != nil {
		logger.Error("bad request", slogx.Error(err))
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(user.Username, user.Password)
	if err != nil {
		logger.Error("failed to generate token", slogx.Error(err))
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Info("user is authorized", slog.String("username", user.Username))
	c.JSON(http.StatusOK, gin.H{"token": token})
}
