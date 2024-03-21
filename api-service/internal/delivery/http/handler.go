package http

import (
	"api-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		post := api.Group("/post")
		{
			post.GET("/:id")
			post.GET("/")
			post.POST("/")

			comment := post.Group("/:id/comment")
			{
				comment.POST("/")
				comment.GET("/")
				comment.GET("/:commentId")
			}
		}

	}

	return router
}
