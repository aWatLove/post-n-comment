package postgres

import (
	"data-service/internal/model"
	"gorm.io/gorm"
)

type PostPostgres struct {
	db *gorm.DB
}

func (p PostPostgres) Create(post model.Post) (int, error) {
	result := p.db.Create(post)
	if result.Error != nil {
		return 0, result.Error
	}
	return post.Id, nil //todo проверить
}

func (p PostPostgres) GetAll() ([]model.Post, error) {
	var posts []model.Post
	result := p.db.Find(&posts)
	return posts, result.Error
}

func (p PostPostgres) GetById(id int) (model.Post, error) {
	var post model.Post
	result := p.db.Where("id=?", id).First(&post)
	return post, result.Error
}

func NewPostPostgres(db *gorm.DB) *PostPostgres {
	return &PostPostgres{db: db}
}
