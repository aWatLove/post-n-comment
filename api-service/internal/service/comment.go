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
	resp, err := http.Get(fmt.Sprintf("%s/api/post/%d/comment", c.dataServiceUrl, id))
	if err != nil {
		log.Printf("error get request: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var comments []model.Comment
	if err = json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		log.Printf("error while decoding response body: %s", err.Error())
		return nil, err
	}

	return comments, nil
}

func (c CommentService) GetCommentById(postId, id int) (model.Comment, error) {
	var comment model.Comment
	resp, err := http.Get(fmt.Sprintf("%s/api/post/%d/comment/%d", c.dataServiceUrl, postId, id))
	if err != nil {
		log.Printf("error get request: %s", err.Error())
		return comment, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("unexpected status code: %d", resp.StatusCode)
		return comment, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&comment); err != nil {
		log.Printf("error while decoding response body: %s", err.Error())
		return comment, err
	}

	return comment, nil
}
