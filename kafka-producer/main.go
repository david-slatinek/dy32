package main

import (
	"context"
	producer "kafka-producer/producer"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	pro := producer.KafkaProducer{
		Address:   "localhost:9092",
		Topic:     "kafka-invoices",
		Partition: 0,
	}

	err := pro.Init()
	if err != nil {
		log.Fatalf("failed to dial leader: %v", err)
	}
	defer func(pro producer.KafkaProducer) {
		if err := pro.Close(); err != nil {
			log.Printf("failed to close writer: %v", err)
		}
	}(pro)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sig := <-signals
		log.Println("Got signal: ", sig)
		log.Println("Signaling other goroutines ...")
		cancel()

		log.Println("Waiting for 5 seconds ...")

		ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel2()

		select {
		case <-ctx2.Done():
			os.Exit(0)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go pro.Write(ctx, &wg)
	wg.Wait()

	//for {
	//	invoice := util.RandomInvoice()
	//	log.Println("Writing invoice with id: ", invoice.ID)
	//
	//	err = pro.Write()
	//	if err != nil {
	//		log.Println("error with write: ", err)
	//	}
	//
	//	time.Sleep(2 * time.Second)
	//}
}
