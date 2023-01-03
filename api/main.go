package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"main/controller"
	"main/producer"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//err := env.Load("env/.env")
	//if err != nil {
	//	log.Fatalf("error with env.Load: %v", err)
	//}
	//
	//uri := os.Getenv("MONGODB_URI")
	//if uri == "" {
	//	log.Fatal("MONGODB_URI is not set")
	//}
	//
	//serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	//clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//
	//client, err := mongo.Connect(ctx, clientOptions)
	//if err != nil {
	//	log.Fatalf("error with Connect: %v", err)
	//}
	//
	//if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
	//	log.Fatalf("error with Ping: %v", err)
	//}

	pro := producer.KafkaProducer{
		Address: "localhost:9092",
		Topic:   "kafka-invoices",
	}

	err := pro.Init()
	if err != nil {
		log.Printf("failed to dial leader: %v", err)
	}
	defer func(p producer.KafkaProducer) {
		if err := p.Close(); err != nil {
			log.Printf("failed to close writer: %v", err)
		}
	}(pro)

	invoiceController := controller.InvoiceController{
		Producer: pro,
	}

	router := gin.Default()
	router.POST("/invoice", invoiceController.Create)

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		log.Println("server is up at: " + srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("ListenAndServe() error: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown() error: %s\n", err)
	}
	log.Println("shutting down")
}
