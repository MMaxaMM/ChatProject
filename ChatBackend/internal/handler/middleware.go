package handler

import (
	"chat"
	"chat/internal/lib/slogx"
	"chat/internal/service"
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

type MiddlewareHandler struct {
	service *service.Service
	log     *slog.Logger
}

func NewMiddlewareHandler(
	service *service.Service,
	log *slog.Logger,
) *MiddlewareHandler {
	return &MiddlewareHandler{service: service, log: log}
}

func (h *MiddlewareHandler) AccessControl(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}

func (h *MiddlewareHandler) UserIdentity(c *gin.Context) {
	const op = "handler.UserIdentity"
	log := h.log.With(slog.String("op", op))

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		log.Error("Empty authorization header")
		NewErrorResponse(c, http.StatusUnauthorized, MsgBadAuthHeader)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		log.Error("Invalid authorization header")
		NewErrorResponse(c, http.StatusUnauthorized, MsgBadAuthHeader)
		return
	}

	if len(headerParts[1]) == 0 {
		log.Error("Authorization token is empty")
		NewErrorResponse(c, http.StatusUnauthorized, MsgBadAuthHeader)
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		if errors.Is(err, chat.ErrTokenExpired) {
			log.Warn("Token expired", slogx.Error(err))
			NewErrorResponse(c, http.StatusUnauthorized, MsgTokenExpired)
			return
		}
		log.Error("Failed to parse token", slogx.Error(err))
		NewErrorResponse(c, http.StatusUnauthorized, MsgBadAuthHeader)
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int64, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int64)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
