package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// Storage представляет собой интерфейс для взаимодействия с хранилищем данных
type Storage interface {
	InsertMessage(date time.Time, key, message string) error
}

// Reader представляет собой интерфейс для чтения сообщений из Kafka
type Reader interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
	Close() error
}

// Consumer представляет собой Kafka consumer, который читает сообщения из Kafka и сохраняет их в базе данных
type Consumer struct {
	reader  Reader
	storage Storage
}

// NewConsumer создает новый экземпляр Kafka consumer
func NewConsumer(reader Reader, storage Storage) *Consumer {
	return &Consumer{
		reader:  reader,
		storage: storage,
	}
}

// Start запускает процесс чтения сообщений из Kafka и сохранения их в базе данных
func (c *Consumer) Start(ctx context.Context) error {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err() // завершение работы по сигналу контекста
			}
			log.Printf("pkg/kafka/consumer.go: could not read message: %v", err)
			continue
		}

		err = c.storage.InsertMessage(time.Now(), string(msg.Key), string(msg.Value))
		if err != nil {
			log.Printf("pkg/kafka/consumer.go: could not insert message: %v", err)
		} else {
			log.Printf("pkg/kafka/consumer.go: message stored: %s", string(msg.Value))
		}
	}
}
