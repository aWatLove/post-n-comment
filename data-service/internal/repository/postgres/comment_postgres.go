package postgres

import (
	"data-service/internal/model"
	"gorm.io/gorm"
)

type CommentPostgres struct {
	db *gorm.DB
}

// todo сделать проверку на наличие поста
func (c CommentPostgres) Create(comment model.Comment) (int, error) {
	result := c.db.Create(&comment)
	if result.Error != nil {
		return 0, result.Error
	}
	return comment.Id, nil //todo проверить
}

func (c CommentPostgres) GetAll(postId int) ([]model.Comment, error) {
	var comments []model.Comment
	result := c.db.Where("post_id=?", postId).Find(&comments)
	return comments, result.Error
}

func (c CommentPostgres) GetById(id int) (model.Comment, error) {
	var comment model.Comment
	result := c.db.Where("id=?", id).First(&comment)
	return comment, result.Error
}

func NewCommentPostgres(db *gorm.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}
