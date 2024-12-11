package handler

import (
	"github.com/gin-gonic/gin"
)

type ErrorMessage string

const (
	MsgInternal            ErrorMessage = "Unknown internal server error"
	MsgBadRequest          ErrorMessage = "Bad request"
	MsgBadAuthHeader       ErrorMessage = "Incorrect authorization header"
	MsgUserExsists         ErrorMessage = "User whit that username already exists"
	MsgUserNotFound        ErrorMessage = "Incorrect username or password"
	MsgServiceNotAvailable ErrorMessage = "Service not available"
	MsgTokenExpired        ErrorMessage = "JWT token expired"
)

type ErrorResponse struct {
	Message ErrorMessage `json:"error"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message ErrorMessage) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}
