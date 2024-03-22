package postgres

import (
	"data-service/internal/model"
	"gorm.io/gorm"
)

type PostPostgres struct {
	db *gorm.DB
}

func (p PostPostgres) GetTopPosts() ([]model.Post, error) {
	var topPosts []model.Post
	err := p.db.
		Select("posts.*, COUNT(comments.id) AS comment_count").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(10).
		Find(&topPosts).Error
	if err != nil {
		return nil, err
	}
	return topPosts, nil
}

func (p PostPostgres) GetAllAuthorsPost(author string) ([]model.Post, error) {
	var posts []model.Post
	err := p.db.Where("author = ?", author).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p PostPostgres) GetTopAuthors() ([]model.TopAuthors, error) {
	var topAuthors []model.TopAuthors
	err := p.db.
		Table("posts").
		Select("author, COUNT(id) AS posts").
		Group("author").
		Order("posts DESC").
		Limit(10).
		Find(&topAuthors).Error
	if err != nil {
		return nil, err
	}
	return topAuthors, nil
}

func (p PostPostgres) Create(post model.Post) (int, error) {
	result := p.db.Create(&post)
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
