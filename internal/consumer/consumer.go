package consumer

import (
	"context"
	"log"
	"logcollector/internal/storage"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader  *kafka.Reader
	storage *storage.Postgres
}

func NewConsumer(brokers []string, topic, groupID string, storage *storage.Postgres) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: groupID,
		}),
		storage: storage,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("internal/consumer/consumer.go: could not read message: %v", err)
			continue
		}

		if err := c.storage.InsertMessage(string(msg.Value)); err != nil {
			log.Printf("internal/consumer/consumer.go: could not insert message: %v", err)
		} else {
			log.Printf("internal/consumer/consumer.go: message stored: %s", string(msg.Value))
		}
	}
}
