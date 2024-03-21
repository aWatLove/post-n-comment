package http

import (
	"api-service/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) CreatePost(c *gin.Context) {
	var input model.Post
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Post.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "post created!")
}

func (h Handler) getAllPosts(c *gin.Context) {

}

func (h Handler) getPostById(c *gin.Context) {

}
