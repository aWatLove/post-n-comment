package model

type TopAuthors struct {
	Author string `json:"author"`
	Posts  int    `json:"posts"`
}

type TopPosts struct {
	Post     Post `json:"post"`
	Comments int  `json:"comments"`
}
