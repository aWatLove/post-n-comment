package kafka

import (
	"context"
	"data-service/internal/model"
	"data-service/internal/service"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type Kafka struct {
	service *service.Service
}

func NewKafka(service *service.Service) *Kafka {
	return &Kafka{service: service}
}

func (k *Kafka) NewConnect(ctx context.Context, network, address, topic string, partition int) (*kafka.Conn, error) {
	con, err := kafka.DialLeader(ctx, network, address, topic, partition)
	if err != nil {
		return nil, err
	}
	return con, nil
}

func (k *Kafka) SubscribePost(conn *kafka.Conn) error {
	for {
		msg, err := conn.ReadMessage(10e6)
		if err != nil {
			log.Printf("kafka: error while reading message: %s", err.Error())
			continue
		}

		post, err := k.UnmarshalPost(msg.Value)
		if err != nil {
			log.Printf("kafka: error while unmarshalling message: %s", err.Error())
			continue
		}

		id, err := k.service.Post.Create(post)
		if err != nil {
			log.Printf("kafka: error while creating post in db: %s", err.Error())
			continue
		}
		log.Printf("added post with id: %d", id)
	}
}

func (k *Kafka) SubscribeComment(conn *kafka.Conn) error {
	for {
		msg, err := conn.ReadMessage(10e6)
		if err != nil {
			log.Printf("kafka: error while reading message: %s", err.Error())
			continue
		}

		comment, err := k.UnmarshalComment(msg.Value)
		if err != nil {
			log.Printf("kafka: error while unmarshalling message: %s", err.Error())
			continue
		}

		id, err := k.service.Comment.Create(comment)
		if err != nil {
			log.Printf("kafka: error while creating comment in db: %s", err.Error())
			continue
		}
		log.Printf("added comment with id: %d", id)
	}
}

func (k *Kafka) UnmarshalPost(msg []byte) (model.Post, error) {
	var post model.Post

	if err := json.Unmarshal(msg, &post); err != nil {
		return post, err
	}
	return post, nil
}

func (k *Kafka) UnmarshalComment(msg []byte) (model.Comment, error) {
	var comment model.Comment

	if err := json.Unmarshal(msg, &comment); err != nil {
		return comment, err
	}
	return comment, nil
}
