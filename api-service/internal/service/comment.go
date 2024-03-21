package service

import (
	"api-service/internal/delivery/kafka"
	"api-service/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CommentService struct {
	kafka          *kafka.Kafka
	dataServiceUrl string
}

func NewCommentService(kafka *kafka.Kafka, dsu string) *CommentService {
	return &CommentService{kafka: kafka, dataServiceUrl: dsu}
}

func (c CommentService) Create(comment model.Comment) error {
	return c.kafka.ProduceComment(comment)
}

func (c CommentService) GetAllComments(id int) ([]model.Comment, error) {
	resp, err := http.Get(fmt.Sprintf("%s/post/%d/comment", c.dataServiceUrl, id))
	if err != nil {
		log.Printf("error get request: %s", err.Error())
		return nil, err
	}
	var data []byte
	_, err = resp.Body.Read(data)
	if err != nil {
		log.Printf("error while reading response body: %s", err.Error())
		return nil, err
	}
	var comments []model.Comment
	err = json.Unmarshal(data, &comments)
	if err != nil {
		log.Printf("error while unmarshaling []comments: %s", err.Error())
		return nil, err
	}
	return comments, nil
}

func (c CommentService) GetCommentById(postId, id int) (model.Comment, error) {
	var comment model.Comment
	resp, err := http.Get(fmt.Sprintf("%s/post/%d/comment/%d", c.dataServiceUrl, postId, id))
	if err != nil {
		log.Printf("error get request: %s", err.Error())
		return comment, err
	}
	var data []byte
	_, err = resp.Body.Read(data)
	if err != nil {
		log.Printf("error while reading response body: %s", err.Error())
		return comment, err
	}

	err = json.Unmarshal(data, &comment)
	if err != nil {
		log.Printf("error while unmarshaling comment: %s", err.Error())
		return comment, err
	}
	return comment, nil
}
