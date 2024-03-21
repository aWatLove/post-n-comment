package kafka

import (
	"api-service/internal/model"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type Kafka struct {
	PostConn    *kafka.Conn
	CommentConn *kafka.Conn
}

func NewKafka(postConn *kafka.Conn, commentConn *kafka.Conn) *Kafka {
	return &Kafka{PostConn: postConn, CommentConn: commentConn}
}

func NewConnect(ctx context.Context, network, address, topic string, partition int) (*kafka.Conn, error) {
	con, err := kafka.DialLeader(ctx, network, address, topic, partition)
	if err != nil {
		log.Printf("failed to dial leader: %s", err.Error())
		return nil, err
	}
	return con, nil
}

func (k *Kafka) ProducePost(post model.Post) error {
	msg, err := k.marshalPost(post)
	if err != nil {
		return err
	}
	_, err = k.PostConn.Write(msg)
	if err != nil {
		log.Printf("failed to write message: %s", err.Error())
		return err
	}
	return nil
}

func (k *Kafka) ProduceComment(comment model.Comment) error {
	msg, err := k.marshalComment(comment)
	if err != nil {
		return err
	}
	_, err = k.CommentConn.Write(msg)
	if err != nil {
		log.Printf("failed to write message: %s", err.Error())
		return err
	}
	return nil
}

func (k Kafka) marshalPost(post model.Post) ([]byte, error) {
	data, err := json.Marshal(post)
	if err != nil {
		log.Printf("error while marshalling post model: %s", err)
		return nil, err
	}
	return data, nil
}

func (k Kafka) marshalComment(comment model.Comment) ([]byte, error) {
	data, err := json.Marshal(comment)
	if err != nil {
		log.Printf("error while marshalling comment model: %s", err)
		return nil, err
	}
	return data, nil
}
