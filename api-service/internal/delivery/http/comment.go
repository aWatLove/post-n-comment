package http

import (
	"api-service/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) CreateComment(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid post id param")
		return
	}
	var input model.Comment
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	input.PostId = postId
	err = h.services.Comment.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "comment created!")
}

func (h Handler) getAllComments(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid post id param")
		return
	}

	comment, err := h.services.Comment.GetAllComments(postId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h Handler) getCommentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //todo доделать проверку на post_id. getCommentById
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid post id param")
		return
	}

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id param")
		return
	}

	post, err := h.services.Comment.GetCommentById(id, commentId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}
