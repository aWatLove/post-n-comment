package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.Post.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h Handler) getPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id param")
		return
	}

	post, err := h.services.Post.GetById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}
