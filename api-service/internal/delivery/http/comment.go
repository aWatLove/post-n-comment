package http

import (
	"api-service/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create
// @Tags comments
// @Description create comment
// @ID create-comment
// @Accept json
// @Produce json
// @Param input body model.CommentRequest true "Comment info"
// @Success 200 {string} string 1
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure default {object} model.ErrorResponse
// @Router /api/comment [post]
func (h Handler) CreateComment(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		model.NewErrorResponse(c, http.StatusBadRequest, "invalid post id param")
		return
	}
	var input model.CommentRequest
	if err := c.BindJSON(&input); err != nil {
		model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	comment := model.Comment{
		PostId: postId,
		Author: input.Author,
		Text:   input.Text,
	}
	err = h.services.Comment.Create(comment)
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "comment created!")
}

// @Summary Get all
// @Tags comments
// @Description get all comments
// @ID getall-comments
// @Accept json
// @Produce json
// @Param id path int true "Post id"
// @Success 200 {object} []model.Comment
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure default {object} model.ErrorResponse
// @Router /api/post/{id}/comment [get]
func (h Handler) getAllComments(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		model.NewErrorResponse(c, http.StatusBadRequest, "invalid post id param")
		return
	}

	comment, err := h.services.Comment.GetAllComments(postId)
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, comment)
}

// @Summary Get by Id
// @Tags comments
// @Description Get comment by id
// @ID getbyid-comment
// @Accept json
// @Produce json
// @Param id path int true "Post id"
// @Param commentId path int true "Comment id"
// @Success 200 {object} model.Post
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure default {object} model.ErrorResponse
// @Router /api/post/{id}/comment/{commentId} [get]
func (h Handler) getCommentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //todo доделать проверку на post_id. getCommentById
	if err != nil {
		model.NewErrorResponse(c, http.StatusBadRequest, "invalid post id param")
		return
	}

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		model.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	post, err := h.services.Comment.GetCommentById(id, commentId)
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}
