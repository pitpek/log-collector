package kafka

import (
	"context"
	logger "log"
	"time"

	"logcollector/internal/schemas"
	"logcollector/internal/service"

	"github.com/segmentio/kafka-go"
)

// Consumer представляет собой Kafka consumer, который читает сообщения из Kafka и сохраняет их в базе данных
type Consumer struct {
	reader  *kafka.Reader
	service *service.Service
}

// NewConsumer создает новый экземпляр Kafka consumer
func NewConsumer(reader *kafka.Reader, service *service.Service) *Consumer {
	return &Consumer{
		reader:  reader,
		service: service,
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
			logger.Printf("pkg/kafka/consumer.go: could not read message: %v", err)
			continue
		}

		log := schemas.Logs{
			Date:    time.Now(),
			AppName: string(msg.Key),
			Message: string(msg.Value),
		}

		err = c.service.AddLog(log)
		if err != nil {
			logger.Printf("pkg/kafka/consumer.go: could not insert message: %v", err)
		} else {
			logger.Printf("pkg/kafka/consumer.go: message stored: %s", string(msg.Value))
		}
	}
}

// Stop останавливает чтение сообщений из Kafka и закрывает reader
func (c *Consumer) Stop() error {
	return c.reader.Close()
}
