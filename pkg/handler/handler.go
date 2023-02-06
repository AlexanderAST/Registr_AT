package handler

import (
	"TBCC_RegistrationPJ/pkg/service"
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
		auth.GET("/sign-in", h.signIn)
	}
	lists := router.Group("/lists", h.userIdentity)
	{

		lists.POST("/login", h.createLogin)
		lists.GET("/count", h.createCount)

	}

	return router
}
