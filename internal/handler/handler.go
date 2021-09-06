package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/seshoo/my-todo/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	h.initApi(router)

	return router
}

func (h *Handler) initApi(router *gin.Engine) {
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sing-up", h.singUp)
			auth.POST("/sing-in", h.SingIn)
		}
	}
}
