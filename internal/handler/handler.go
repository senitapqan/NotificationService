package handler

import (
	"goNotificationService/dtos"
	"goNotificationService/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Handler struct {
	service  *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service:   service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.POST("/send", h.sendMessage)

	return router
}

func (h *Handler) sendMessage(c *gin.Context) {
	var input dtos.SendMessageRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "wrong format of request")
		return
	}

	err := h.service.SendMessage(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "something went wrong: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message":"all emails are delivered",
	})
}