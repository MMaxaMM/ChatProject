package handler

import (
	"chat/internal/service"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

type Auth interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type Middleware interface {
	AccessControl(c *gin.Context)
	UserIdentity(c *gin.Context)
}

type Control interface {
	GetStart(c *gin.Context)
	CreateChat(c *gin.Context)
	DeleteChat(c *gin.Context)
	GetHistory(c *gin.Context)
}

type Chat interface {
	SendMessage(c *gin.Context)
}

type RAG interface {
	SendMessageRAG(c *gin.Context)
}

type Audio interface {
	Recognize(c *gin.Context)
}

type Video interface {
	Detect(c *gin.Context)
}

type Handler struct {
	Auth
	Middleware
	Control
	Chat
	RAG
	Audio
	Video
}

func NewHandler(service *service.Service, log *slog.Logger) *Handler {
	return &Handler{
		Auth:       NewAuthHandler(service, log),
		Middleware: NewMiddlewareHandler(service, log),
		Control:    NewControlHandler(service, log),
		Chat:       NewChatHandler(service, log),
		RAG:        NewRAGHandler(service, log),
		Audio:      NewAudioHandler(service, log),
		Video:      NewVideoHandler(service, log),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth", h.AccessControl)
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	control := router.Group(
		"/control",
		h.AccessControl,
		h.UserIdentity,
	)
	{
		control.GET("/start", h.GetStart)
		control.POST("/create", h.CreateChat)
		control.DELETE("/delete", h.DeleteChat)
		control.POST("/history", h.GetHistory)

	}

	chat := router.Group(
		"/chat",
		h.AccessControl,
		h.UserIdentity,
	)
	{
		chat.POST("/message", h.SendMessage)
	}

	rag := router.Group(
		"/rag",
		h.AccessControl,
		h.UserIdentity,
	)
	{
		rag.POST("/message", h.SendMessageRAG)
	}

	audio := router.Group(
		"/audio",
		h.AccessControl,
		h.UserIdentity,
	)
	{
		audio.POST("/recognize", h.Recognize)
	}

	video := router.Group(
		"/video",
		h.AccessControl,
		h.UserIdentity,
	)
	{
		video.POST("/detect", h.Detect)
	}

	return router
}
