package model

import "time"

type Comment struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	PostId    int       `json:"post_id" gorm:"foreignKey:Post"`
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"CreatedAt" format:"2021-11-26T06:22:19Z"`
}
