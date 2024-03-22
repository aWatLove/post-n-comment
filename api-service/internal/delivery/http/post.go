package http

import (
	"api-service/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create
// @Tags posts
// @Description create post
// @ID create-post
// @Accept json
// @Produce json
// @Param input body model.PostRequest true "Post info"
// @Success 200 {string} string 1
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure default {object} model.ErrorResponse
// @Router /api/post [post]
func (h Handler) CreatePost(c *gin.Context) {
	var input model.PostRequest
	if err := c.BindJSON(&input); err != nil {
		model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	post := model.Post{Text: input.Text, Author: input.Author}
	err := h.services.Post.Create(post)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "post created!")
}

// @Summary Get all
// @Tags posts
// @Description get all posts
// @ID getall-posts
// @Accept json
// @Produce json
// @Success 200 {object} []model.Post
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure default {object} model.ErrorResponse
// @Router /api/post [get]
func (h Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.Post.GetAllPosts()
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

// @Summary Get by Id
// @Tags posts
// @Description Get post by id
// @ID getbyid-post
// @Accept json
// @Produce json
// @Param id path int true "Post id"
// @Success 200 {object} model.Post
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure default {object} model.ErrorResponse
// @Router /api/post/{id} [get]
func (h Handler) getPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		model.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	post, err := h.services.Post.GetPostById(id)
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

// @Summary Get top
// @Tags posts
// @Description get top posts
// @ID gettop-posts
// @Accept json
// @Produce json
// @Success 200 {object} []model.Post
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure default {object} model.ErrorResponse
// @Router /api/post/top [get]
func (h Handler) getTopPosts(c *gin.Context) {
	posts, err := h.services.Post.GetTopPosts()
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}

// @Summary Get top
// @Tags authors
// @Description get top authors
// @ID gettop-authors
// @Accept json
// @Produce json
// @Success 200 {object} []model.TopAuthors
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure default {object} model.ErrorResponse
// @Router /api/author/top [get]
func (h Handler) getTopAuthors(c *gin.Context) {
	topAuthors, err := h.services.Post.GetTopAuthors()
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, topAuthors)
}

// @Summary Get by Author
// @Tags authors
// @Description Get author's posts
// @ID getbyauthor-post
// @Accept json
// @Produce json
// @Param author path string true "Author name"
// @Success 200 {object} []model.Post
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Failure default {object} model.ErrorResponse
// @Router /api/author/post/{author} [get]
func (h Handler) getAuthorsPosts(c *gin.Context) {
	authorName := c.Param("author")
	posts, err := h.services.Post.GetAllAuthorsPost(authorName)
	if err != nil {
		model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, posts)
}
