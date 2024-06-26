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

func (h Handler) getTopPosts(c *gin.Context) {
	posts, err := h.services.Post.GetTopPosts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h Handler) getTopAuthors(c *gin.Context) {
	topAuthors, err := h.services.Post.GetTopAuthors()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, topAuthors)
}

func (h Handler) getAuthorsPosts(c *gin.Context) {
	authorName := c.Param("author")
	posts, err := h.services.Post.GetAllAuthorsPost(authorName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}
