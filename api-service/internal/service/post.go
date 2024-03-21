package service

import (
	"api-service/internal/delivery/kafka"
	"api-service/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PostService struct {
	kafka          *kafka.Kafka
	dataServiceUrl string
}

func NewPostService(kafka *kafka.Kafka, dsu string) *PostService {
	return &PostService{kafka: kafka, dataServiceUrl: dsu}
}

func (p PostService) Create(post model.Post) error {
	return p.kafka.ProducePost(post)
}

func (p PostService) GetAllPosts() ([]model.Post, error) {
	resp, err := http.Get(fmt.Sprintf("%s/post", p.dataServiceUrl))
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
	var posts []model.Post
	err = json.Unmarshal(data, &posts)
	if err != nil {
		log.Printf("error while unmarshaling []post: %s", err.Error())
		return nil, err
	}
	return posts, nil
}

func (p PostService) GetPostById(id int) (model.Post, error) {
	var post model.Post
	resp, err := http.Get(fmt.Sprintf("%s/post/%d", p.dataServiceUrl, id))
	if err != nil {
		log.Printf("error get request: %s", err.Error())
		return post, err
	}
	var data []byte
	_, err = resp.Body.Read(data)
	if err != nil {
		log.Printf("error while reading response body: %s", err.Error())
		return post, err
	}

	err = json.Unmarshal(data, &post)
	if err != nil {
		log.Printf("error while unmarshaling post: %s", err.Error())
		return post, err
	}
	return post, nil
}
