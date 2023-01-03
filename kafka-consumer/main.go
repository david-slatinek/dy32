package main

import (
	"fmt"
	"log"
	"main/consumer"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	con := consumer.KafkaConsumer{
		Address: "localhost:9092",
		Topic:   "kafka-invoices",
	}

	err := con.Init()
	if err != nil {
		log.Fatalf("failed to dial leader: %v", err)
	}
	defer func(con consumer.KafkaConsumer) {
		if err := con.Close(); err != nil {
			log.Printf("failed to close reader: %v", err)
		}
	}(con)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	go func() {
		sig := <-signals
		log.Printf("got signal: %v", sig)
		os.Exit(0)
	}()

	for {
		inv, err := con.Read()
		if err != nil {
			log.Printf("error with read: %v", err)
		}
		fmt.Println(inv)
	}
}
