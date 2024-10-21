package handler

import (
	"chat/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	chatInterface := router.Group("/chat", h.userIdentity)
	{
		chatInterface.POST("/history", h.GetHistory)
		chatInterface.POST("/message", h.SendMessage)
		chatInterface.POST("/delete", h.Delete)
	}

	return router
}
