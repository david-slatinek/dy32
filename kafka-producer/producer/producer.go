package kafkaProducer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"kafka-producer/util"
	"log"
	"sync"
	"time"
)

type KafkaProducer struct {
	Address   string
	Topic     string
	Partition int
	writer    *kafka.Writer
}

func (receiver *KafkaProducer) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := kafka.DialLeader(ctx, "tcp", receiver.Address, receiver.Topic, receiver.Partition)
	if err != nil {
		return err
	}

	receiver.writer = &kafka.Writer{
		Addr:         kafka.TCP(receiver.Address),
		Topic:        receiver.Topic,
		WriteTimeout: 4 * time.Second,
	}
	return nil
}

func (receiver *KafkaProducer) Close() error {
	return receiver.writer.Close()
}

func (receiver *KafkaProducer) Write(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			break
		}

		err := receiver.writeRandom()
		if err != nil {
			log.Println("error with write: ", err)
		}
		time.Sleep(2 * time.Second)
	}
}

func (receiver *KafkaProducer) writeRandom() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	invoice := util.RandomInvoice()
	log.Println("Writing invoice with id: ", invoice.ID)

	inv, err := invoice.Json()
	if err != nil {
		return err
	}

	err = receiver.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(invoice.ID),
		Value: inv,
	})
	return err
}
