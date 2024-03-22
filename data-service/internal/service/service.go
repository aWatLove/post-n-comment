package service

import (
	"data-service/internal/model"
	"data-service/internal/repository"
)

type Post interface {
	Create(post model.Post) (int, error)
	GetAll() ([]model.Post, error)
	GetById(id int) (model.Post, error)
	GetTopPosts() ([]model.Post, error)
	GetAllAuthorsPost(string) ([]model.Post, error)
	GetTopAuthors() ([]model.TopAuthors, error)
}

type Comment interface {
	Create(model.Comment) (id int, err error)
	GetAll(postId int) ([]model.Comment, error)
	GetById(id int) (model.Comment, error)
}

type Service struct {
	Post
	Comment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Post:    NewPostService(repos.Post),
		Comment: NewCommentService(repos.Comment),
	}
}
