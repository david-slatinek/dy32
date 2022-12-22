package kafkaProducer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"kafka-producer/random"
	"log"
	"sync"
	"time"
)

type KafkaProducer struct {
	Address string
	Topic   string
	Delay   time.Duration
	writer  *kafka.Writer
}

func (receiver *KafkaProducer) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := kafka.DialLeader(ctx, "tcp", receiver.Address, receiver.Topic, 0)
	if err != nil {
		return err
	}

	receiver.writer = &kafka.Writer{
		Addr:         kafka.TCP(receiver.Address),
		Topic:        receiver.Topic,
		WriteTimeout: 5 * time.Second,
	}
	return nil
}

func (receiver *KafkaProducer) Close() error {
	return receiver.writer.Close()
}

func (receiver *KafkaProducer) Write(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	t := time.Now()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if time.Now().After(t.Add(receiver.Delay)) {
				t = time.Now()
				err := receiver.writeRandom()
				if err != nil {
					log.Printf("error with write: %v", err)
				}
			}
			break
		}
	}
}

func (receiver *KafkaProducer) writeRandom() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	invoice := random.Invoice()
	log.Printf("writing invoice with id: %s", invoice.ID)

	inv, err := invoice.Json()
	if err != nil {
		return err
	}

	return receiver.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(invoice.ID),
		Value: inv,
	})
}
