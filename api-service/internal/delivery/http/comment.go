package http

import (
	"api-service/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) CreateComment(c *gin.Context) {
	var input model.Comment
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Comment.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "comment created!")
}

func (h Handler) getAllComments(c *gin.Context) {

}

func (h Handler) getCommentById(c *gin.Context) {

}
