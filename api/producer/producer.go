package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"main/model"
	"time"
)

type KafkaProducer struct {
	Address string
	Topic   string
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

func (receiver *KafkaProducer) Write(invoice model.Invoice) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	inv, err := invoice.ToJson()
	if err != nil {
		return err
	}

	return receiver.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(invoice.ID),
		Value: inv,
	})
}
