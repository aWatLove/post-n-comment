package service

import (
	"data-service/internal/model"
	"data-service/internal/repository"
)

type PostService struct {
	repo repository.Post
}

func (p PostService) Create(post model.Post) (int, error) {
	return p.repo.Create(post)
}

func (p PostService) GetAll() ([]model.Post, error) {
	return p.repo.GetAll()
}

func (p PostService) GetById(id int) (model.Post, error) {
	return p.repo.GetById(id)
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}
