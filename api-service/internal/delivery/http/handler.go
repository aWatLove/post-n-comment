package http

import (
	_ "api-service/docs"
	"api-service/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
		author := api.Group("/author")
		{
			author.GET("/top", h.getTopAuthors)
			author.GET("/post/:author", h.getAuthorsPosts)
			// author.GET("/comment/:author") // все комментарии автора
		}
		post := api.Group("/post")
		{
			post.GET("/:id", h.getPostById)
			post.GET("/", h.getAllPosts)
			post.GET("/top", h.getTopPosts)
			post.POST("/", h.CreatePost)

			comment := post.Group("/:id/comment")
			{
				comment.POST("/", h.CreateComment)
				comment.GET("/", h.getAllComments)
				comment.GET("/:commentId", h.getCommentById)
			}
		}

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
