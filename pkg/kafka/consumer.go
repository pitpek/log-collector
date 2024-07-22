package kafka

import (
	"context"
	"log"
	"time"

	"logcollector/internal/repository"

	"github.com/segmentio/kafka-go"
)

// Consumer представляет собой Kafka consumer, который читает сообщения из Kafka и сохраняет их в базе данных
type Consumer struct {
	reader *kafka.Reader
	repo   *repository.Repository
}

// NewConsumer создает новый экземпляр Kafka consumer
func NewConsumer(reader *kafka.Reader, repo *repository.Repository) *Consumer {
	return &Consumer{
		reader: reader,
		repo:   repo,
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

		err = c.repo.AddLog(time.Now(), string(msg.Key), string(msg.Value))
		if err != nil {
			log.Printf("pkg/kafka/consumer.go: could not insert message: %v", err)
		} else {
			log.Printf("pkg/kafka/consumer.go: message stored: %s", string(msg.Value))
		}
	}
}

// Stop останавливает чтение сообщений из Kafka и закрывает reader
func (c *Consumer) Stop() error {
	return c.reader.Close()
}
