package service

import (
	"api-service/internal/delivery/kafka"
	"api-service/internal/model"
)

type Post interface {
	Create(model.Post) error
	GetAllPosts() ([]model.Post, error)
	GetPostById(int) (model.Post, error)
	GetTopPosts() ([]model.Post, error)
	GetAllAuthorsPost(string) ([]model.Post, error)
	GetTopAuthors() ([]model.TopAuthors, error)
}

type Comment interface {
	Create(model.Comment) error
	GetAllComments(int) ([]model.Comment, error)
	GetCommentById(postId int, id int) (model.Comment, error)
}

type Service struct {
	Post
	Comment
}

func NewService(kafka *kafka.Kafka, dsu string) *Service {
	return &Service{
		Post:    NewPostService(kafka, dsu),
		Comment: NewCommentService(kafka, dsu),
	}
}
