package app

import "github.com/gin-gonic/gin"

type ControlHandlerInterface interface {
	GetStart(c *gin.Context)
	CreateChat(c *gin.Context)
	DeleteChat(c *gin.Context)
}

type MiddlewareHandlerInterface interface {
	AccessControl(c *gin.Context)
	UserIdentity(c *gin.Context)
}

type AuthHandlerInterface interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type ChatHandlerInterface interface {
	GetHistory(c *gin.Context)
	SendMessage(c *gin.Context)
}

type AudioHandlerInterface interface {
	GetHistory(c *gin.Context)
	Recognize(c *gin.Context)
}

type VideoHandlerInterface interface {
	GetHistory(c *gin.Context)
	Detect(c *gin.Context)
}
