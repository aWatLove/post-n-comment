package service

import (
	"data-service/internal/model"
	"data-service/internal/repository"
)

type CommentService struct {
	repo repository.Comment
}

func (c CommentService) Create(comment model.Comment) (id int, err error) {
	return c.repo.Create(comment)
}

func (c CommentService) GetAll(postId int) ([]model.Comment, error) {
	return c.repo.GetAll(postId)
}

func (c CommentService) GetById(id int) (model.Comment, error) {
	return c.repo.GetById(id)
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}
