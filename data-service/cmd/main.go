package main

import (
	"context"
	"data-service/internal/delivery/http"
	"data-service/internal/delivery/kafka"
	"data-service/internal/model"
	"data-service/internal/repository"
	"data-service/internal/repository/postgres"
	"data-service/internal/service"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error while initializing .env variables: %s", err.Error())
	}

	db, err := postgres.ConnectDB(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("error while connecting db: %s", err.Error())
	}

	err = db.AutoMigrate(model.Post{}, model.Comment{})
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	//kafka init
	var wg sync.WaitGroup
	k := kafka.NewKafka(services)
	kAddress := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_ADDRESS"), os.Getenv("KAFKA_PORT"))
	readerPost := k.NewReaderConn(kAddress, "main-group", os.Getenv("KAFKA_TOPIC_POST"))
	readerComment := k.NewReaderConn(kAddress, "main-group", os.Getenv("KAFKA_TOPIC_COMMENT"))

	// post topic connection and consume
	wg.Add(1)
	go func() {
		if err = k.SubscribePost(readerPost); err != nil {
			log.Fatalf("error while consume post topic: %s", err.Error())
		}
	}()

	// comment topic connection and consume
	wg.Add(1)
	go func() {
		if err = k.SubscribeComment(readerComment); err != nil {
			log.Fatalf("error while consume comment topic: %s", err.Error())
		}
	}()

	//server init
	handler := http.NewHandler(services)
	srv := new(http.Server)

	go func() {
		if err = srv.Run(os.Getenv("APP_PORT"), handler.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("server started")

	//graceful shuttdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("server shutting down")
	if err = srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	cdb, _ := db.DB()
	if err = cdb.Close(); err != nil {
		log.Printf("error while closing db connection: %s", err.Error())
	}

	if err = readerPost.Close(); err != nil {
		log.Printf("error while closing kafka post topic connection: %s", err.Error())
	}
	wg.Done()

	if err = readerComment.Close(); err != nil {
		log.Printf("error while closing kafka comment topic connection: %s", err.Error())
	}
	wg.Done()

	wg.Wait()
}
