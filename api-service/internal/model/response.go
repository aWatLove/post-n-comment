package model

import (
	"github.com/gin-gonic/gin"
	"log"
)

type TopAuthors struct {
	Author string `json:"author"`
	Posts  int    `json:"posts"`
}

type TopPosts struct {
	Post     Post `json:"post"`
	Comments int  `json:"comments"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Printf(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}
