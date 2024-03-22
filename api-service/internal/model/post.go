package model

import "time"

type Post struct {
	Id        int       `json:"id"`
	Text      string    `json:"text"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"CreatedAt" format:"2021-11-26T06:22:19Z"`
}
