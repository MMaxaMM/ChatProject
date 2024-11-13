package app

import (
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type Handler struct {
	middlewareHandler MiddlewareHandlerInterface
	authHandler       AuthHandlerInterface
	controlHandler    ControlHandlerInterface
	chatHandler       ChatHandlerInterface
	audioHandler      AudioHandlerInterface
	videoHandler      VideoHandlerInterface
}

func NewHandler(
	middlewareHandler MiddlewareHandlerInterface,
	authHandler AuthHandlerInterface,
	controlHandler ControlHandlerInterface,
	chatHandler ChatHandlerInterface,
	audioHandler AudioHandlerInterface,
	videoHandler VideoHandlerInterface,
) *Handler {
	return &Handler{
		middlewareHandler: middlewareHandler,
		authHandler:       authHandler,
		controlHandler:    controlHandler,
		chatHandler:       chatHandler,
		audioHandler:      audioHandler,
		videoHandler:      videoHandler,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth", h.middlewareHandler.AccessControl)
	{
		auth.POST("/sign-up", h.authHandler.SignUp)
		auth.POST("/sign-in", h.authHandler.SignIn)
	}

	control := router.Group(
		"/control",
		h.middlewareHandler.AccessControl,
		h.middlewareHandler.UserIdentity,
	)
	{
		control.GET("/start", h.controlHandler.GetStart)
		control.POST("/create", h.controlHandler.CreateChat)
		control.DELETE("/delete", h.controlHandler.DeleteChat)
	}

	chat := router.Group(
		"/chat",
		h.middlewareHandler.AccessControl,
		h.middlewareHandler.UserIdentity,
	)
	{
		chat.POST("/history", h.chatHandler.GetHistory)
		chat.POST("/message", h.chatHandler.SendMessage)
	}

	audio := router.Group(
		"/audio",
		h.middlewareHandler.AccessControl,
		h.middlewareHandler.UserIdentity,
	)
	{
		audio.POST("/history", h.audioHandler.GetHistory)
		audio.POST("/recognize", h.audioHandler.Recognize)
	}

	video := router.Group(
		"/video",
		h.middlewareHandler.AccessControl,
		h.middlewareHandler.UserIdentity,
	)
	{
		video.POST("/history", h.videoHandler.GetHistory)
		video.POST("/detect", h.videoHandler.Detect)
	}

	// TODO: rag

	return router
}
