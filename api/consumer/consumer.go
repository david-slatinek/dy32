package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"main/model"
	"time"
)

type KafkaConsumer struct {
	Address string
	Topic   string
	reader  *kafka.Reader
}

func (receiver *KafkaConsumer) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := kafka.DialLeader(ctx, "tcp", receiver.Address, receiver.Topic, 0)
	if err != nil {
		return err
	}

	receiver.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{receiver.Address},
		Topic:     receiver.Topic,
		Partition: 0,
	})
	return nil
}

func (receiver *KafkaConsumer) Close() error {
	return receiver.reader.Close()
}

func (receiver *KafkaConsumer) Read() (model.Invoice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m, err := receiver.reader.ReadMessage(ctx)
	if err != nil {
		return model.Invoice{}, err
	}
	return model.FromJson(m.Value)
}
