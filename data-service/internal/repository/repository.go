package repository

import (
	"data-service/internal/model"
	"data-service/internal/repository/postgres"
	"gorm.io/gorm"
)

type Post interface {
	Create(post model.Post) (int, error)
	GetAll() ([]model.Post, error)
	GetById(id int) (model.Post, error)
}

type Comment interface {
	Create(model.Comment) (id int, err error)
	GetAll(postId int) ([]model.Comment, error)
	GetById(id int) (model.Comment, error)
}

type Repository struct {
	Post
	Comment
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Post:    postgres.NewPostPostgres(db),
		Comment: postgres.NewCommentPostgres(db),
	}
}
