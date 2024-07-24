package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// Producer представляет собой структуру для отправки сообщений в Kafka.
type Producer struct {
	writer *kafka.Writer
}

// NewProducer создает новый экземпляр Kafka producer.
// writer - экземпляр kafka.Writer, который будет использоваться для отправки сообщений.
func NewProducer(writer *kafka.Writer) *Producer {
	return &Producer{
		writer: writer,
	}
}

// Start запускает процесс отправки сообщений в Kafka.
// ctx - контекст для управления жизненным циклом процесса.
// key - ключ сообщения для Kafka.
func (p *Producer) Start(ctx context.Context, key string) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			// Здесь можно добавить логику для получения данных из ClickHouse или другого источника
			messageValue := "some message" // Замените это на реальные данные

			err := p.writer.WriteMessages(ctx, kafka.Message{
				Key:   []byte(key),
				Value: []byte(messageValue),
			})
			if err != nil {
				log.Printf("pkg/kafka/producer.go: could not write message: %v", err)
			} else {
				log.Printf("pkg/kafka/producer.go: message sent: %s", messageValue)
			}
		}
	}
}
