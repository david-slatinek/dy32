package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"main/consumer"
	"main/controller"
	"main/db"
	"main/env"
	"main/producer"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const address = "localhost:9092"
const topic = "kafka-invoices"

func main() {
	pro := producer.KafkaProducer{
		Address: address,
		Topic:   topic,
	}

	err := pro.Init()
	if err != nil {
		log.Fatalf("failed to dial leader: %v", err)
	}
	defer func(p producer.KafkaProducer) {
		if err := p.Close(); err != nil {
			log.Printf("failed to close writer: %v", err)
		}
	}(pro)

	reader := consumer.KafkaConsumer{
		Address: address,
		Topic:   topic,
	}

	err = reader.Init()
	if err != nil {
		log.Fatalf("failed to dial leader: %v", err)
	}
	defer func(c consumer.KafkaConsumer) {
		if err := c.Close(); err != nil {
			log.Printf("failed to close reader: %v", err)
		}
	}(reader)

	err = env.Load("env/.env")
	if err != nil {
		log.Fatalf("error with env.Load: %v", err)
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("error with Connect: %v", err)
	}
	defer func(client *mongo.Client) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := client.Disconnect(ctx); err != nil {
			log.Printf("error with Disconnect: %v", err)
		}
	}(client)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("error with Ping: %v", err)
	}

	invoiceController := controller.InvoiceController{
		Producer: pro,
		Consumer: reader,
		Collection: db.InvoiceCollection{
			Collection: client.Database("invoices-db").Collection("invoices"),
		},
	}

	router := gin.Default()
	kafka := router.Group("/kafka")
	{
		kafka.POST("/invoice", invoiceController.Create)
		kafka.GET("/invoice", invoiceController.Read)
	}
	mongoGroup := router.Group("/mongo")
	{
		mongoGroup.POST("/invoice", invoiceController.CreateMongo)
	}

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("ListenAndServe() error: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown() error: %s\n", err)
	}
	log.Println("shutting down")
}
