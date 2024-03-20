package http

import (
	"data-service/internal/service"
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
			post.GET("/:id", h.getPostById)
			post.GET("/", h.getAllPosts)

			comment := post.Group("/:id/comment")
			{
				comment.GET("/", h.getAllComments)
				comment.GET("/:commentId", h.getCommentById)
			}
		}

	}

	return router
}
