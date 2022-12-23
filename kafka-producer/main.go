package main

import (
	"context"
	producer "kafka-producer/producer"
	"kafka-producer/random"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const address = "localhost:9092"

var producers = []producer.KafkaProducer{
	{
		Address: address,
		Topic:   "kafka-invoices",
		Delay:   2 * time.Second,
		Random:  random.Invoice,
	},
	{
		Address: address,
		Topic:   "kafka-customer",
		Delay:   5 * time.Second,
		Random:  random.Customer,
	},
	{
		Address: address,
		Topic:   "kafka-purchase",
		Delay:   1 * time.Second,
		Random:  random.Purchase,
	},
	{
		Address: address,
		Topic:   "kafka-equipment",
		Delay:   3 * time.Second,
		Random:  random.Equipment,
	},
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for k, v := range producers {
		p := v
		err := p.Init()
		if err != nil {
			log.Printf("index=%d, failed to dial leader: %v", k, err)
			continue
		}
		defer func(p producer.KafkaProducer) {
			if err := p.Close(); err != nil {
				log.Printf("failed to close writer: %v", err)
			}
		}(p)
		wg.Add(1)
		go p.Write(ctx, &wg)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	go func() {
		sig := <-signals
		log.Printf("got signal: %v", sig)
		log.Println("signaling other goroutines ...")
		cancel()

		log.Println("waiting for 5 seconds ...")

		ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel2()

		select {
		case <-ctx2.Done():
			log.Println("exiting")
			os.Exit(0)
		}
	}()

	wg.Wait()
}
