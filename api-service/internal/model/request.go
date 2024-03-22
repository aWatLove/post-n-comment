package model

type PostRequest struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type CommentRequest struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}
