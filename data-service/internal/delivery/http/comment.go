package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) getAllComments(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid post id param")
		return
	}

	comment, err := h.services.Comment.GetAll(postId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h Handler) getCommentById(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id")) //todo доделать проверку на post_id. getCommentById
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid post id param")
		return
	}

	id, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id param")
		return
	}

	post, err := h.services.Comment.GetById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}
