package consumer

import (
	"context"
	"log"
	"logcollector/internal/storage"

	"github.com/segmentio/kafka-go"
)

// Consumer представляет собой Kafka consumer, который читает сообщения из Kafka и сохраняет их в базе данных PostgreSQL
type Consumer struct {
	reader  *kafka.Reader
	storage *storage.Postgres
}

// NewConsumer создает новый экземпляр Kafka consumer
// brokers - список брокеров Kafka
// topic - тема Kafka, из которой нужно читать сообщения
// groupID - ID группы потребителей
// storage - экземпляр PostgreSQL для сохранения сообщений
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

// Start запускает процесс чтения сообщений из Kafka и сохранения их в базе данных
// ctx - контекст для управления жизненным циклом процесса
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
