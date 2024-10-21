package handler

import (
	"chat"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetHistory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request := new(chat.HistoryRequest)
	request.UserId = userId
	if err := c.BindJSON(request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	response, err := h.services.ChatInterface.GetHistory(request)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) SendMessage(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request := new(chat.ChatItem)
	request.UserId = userId
	if err := c.BindJSON(request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	response, err := h.services.ChatInterface.SendMessage(request)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) Delete(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	request := new(chat.HistoryRequest)
	request.UserId = userId
	if err := c.BindJSON(request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.services.ChatInterface.DeleteChat(request)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
