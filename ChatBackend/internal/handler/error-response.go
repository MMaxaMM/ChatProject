package handler

import (
	"github.com/gin-gonic/gin"
)

type ErrorMessage string

const (
	MsgInternal            ErrorMessage = "unknown internal server error"
	MsgBadRequest          ErrorMessage = "bad request"
	MsgBadAuthHeader       ErrorMessage = "incorrect authorization header"
	MsgUserExsists         ErrorMessage = "user whit that username already exists"
	MsgUserNotFound        ErrorMessage = "incorrect username or password"
	MsgServiceNotAvailable ErrorMessage = "service not available"
)

type ErrorResponse struct {
	Message ErrorMessage `json:"error"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message ErrorMessage) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}
