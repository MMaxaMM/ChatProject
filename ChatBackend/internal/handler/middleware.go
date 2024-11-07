package handler

import (
	"chat/internal/lib/slogx"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) accessControl(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}

func (h *Handler) userIdentity(c *gin.Context) {
	const op = "handler.userIdentity"
	logger := h.logger.With(slog.String("op", op))

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		logger.Error("empty auth header")
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		logger.Error("invalid auth header")
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		logger.Error("token is empty")
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		logger.Error("failed to parse token", slogx.Error(err))
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
