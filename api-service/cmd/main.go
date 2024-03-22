package main

import (
	"api-service/internal/delivery/http"
	"api-service/internal/delivery/kafka"
	"api-service/internal/service"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title API-SERVICE
// @version 1.0
// @contact.name Suvorov Vladislav

// @host localhost:8080
// @BasePath /

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error while initializing .env variables: %s", err.Error())
	}

	// kafka
	kAddress := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_ADDRESS"), os.Getenv("KAFKA_PORT"))
	kNetwork := os.Getenv("KAFKA_NETWORK")
	postConn, err := kafka.NewConnect(context.Background(), kNetwork, kAddress, os.Getenv("KAFKA_TOPIC_POST"), 0)
	if err != nil {
		log.Fatal(err)
	}
	commentConn, err := kafka.NewConnect(context.Background(), kNetwork, kAddress, os.Getenv("KAFKA_TOPIC_COMMENT"), 0)
	if err != nil {
		log.Fatal(err)
	}
	k := kafka.NewKafka(postConn, commentConn)
	// services init
	services := service.NewService(k, fmt.Sprintf("%s:%s", os.Getenv("DATASERVICE_ADDRESS"), os.Getenv("DATASERVICE_PORT")))
	// handlers init
	handler := http.NewHandler(services)

	// server init
	srv := new(http.Server)

	go func() {
		if err = srv.Run(os.Getenv("APP_PORT"), handler.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("server shutting down")
	if err = srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	if err = postConn.Close(); err != nil {
		log.Printf("error while closing kafka post topic connection: %s", err.Error())
	}

	if err = commentConn.Close(); err != nil {
		log.Printf("error while closing kafka comment topic connection: %s", err.Error())
	}
}
