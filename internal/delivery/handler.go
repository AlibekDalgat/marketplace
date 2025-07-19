package delivery

import (
	"github.com/gin-gonic/gin"
	"marketplace/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		adverts := api.Group("/adverts")
		{
			adverts.POST("/", h.userIdentity(true), h.createAdvert)
			adverts.GET("/", h.userIdentity(false), h.listAdverts)
		}
	}
	return router
}
