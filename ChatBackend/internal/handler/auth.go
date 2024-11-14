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

type AuthHandler struct {
	service *service.Service
	log     *slog.Logger
}

func NewAuthHandler(service *service.Service, log *slog.Logger) *AuthHandler {
	return &AuthHandler{service: service, log: log}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	const op = "handler.SignUp"
	log := h.log.With(slog.String("op", op))

	request := new(models.SignUpRequest)

	if err := c.BindJSON(request); err != nil {
		log.Error("Bad request", slogx.Error(err))
		NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
		return
	}

	log = log.With(slog.String("username", request.Username))

	response, err := h.service.CreateUser(request)
	if err != nil {
		if errors.Is(err, chat.ErrUserDuplicate) {
			log.Warn("User already exists")
			NewErrorResponse(c, http.StatusConflict, MsgUserExsists)
			return
		}
		log.Error("Failed to create user", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return
	}

	log.Info("New user created")

	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	const op = "handler.SignIn"
	log := h.log.With(slog.String("op", op))

	request := new(models.SignInRequest)

	if err := c.BindJSON(&request); err != nil {
		log.Error("Bad request", slogx.Error(err))
		NewErrorResponse(c, http.StatusBadRequest, MsgBadRequest)
		return
	}

	log = log.With(slog.String("username", request.Username))

	response, err := h.service.GenerateToken(request)
	if err != nil {
		if errors.Is(err, chat.ErrUserNotFound) {
			log.Warn("Incorrect username or password", slogx.Error(err))
			NewErrorResponse(c, http.StatusUnauthorized, MsgUserNotFound)
			return
		}
		log.Error("Failed to generate token", slogx.Error(err))
		NewErrorResponse(c, http.StatusInternalServerError, MsgInternal)
		return

	}

	log.Info("User is authorized")

	c.JSON(http.StatusOK, response)
}
