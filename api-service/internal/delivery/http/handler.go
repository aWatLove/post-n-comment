package http

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
