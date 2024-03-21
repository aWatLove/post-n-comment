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

type dataPost struct {
	data []model.Post
}

func (p PostService) GetAllPosts() ([]model.Post, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/post", p.dataServiceUrl))
	if err != nil {
		log.Printf("error get request: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var posts []model.Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		log.Printf("error while decoding response body: %s", err.Error())
		return nil, err
	}

	return posts, nil

}

func (p PostService) GetPostById(id int) (model.Post, error) {
	var post model.Post
	resp, err := http.Get(fmt.Sprintf("%s/api/post/%d", p.dataServiceUrl, id))
	if err != nil {
		log.Printf("error get request: %s", err.Error())
		return post, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("unexpected status code: %d", resp.StatusCode)
		return post, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		log.Printf("error while decoding response body: %s", err.Error())
		return post, err
	}

	return post, nil

}
