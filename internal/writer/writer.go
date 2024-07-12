package producer

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Writer struct {
	writer *kafka.Writer
}

// NewProducer создает новый экземпляр Kafka producer
// brokers - список брокеров Kafka
// topic - тема Kafka, в которую нужно отправлять сообщения
func NewWriter(brokers []string, topic string) *Writer {
	return &Writer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

// Start запускает процесс отправки сообщений в Kafka
// ctx - контекст для управления жизненным циклом процесса
func (p *Writer) Start(ctx context.Context, key string) error {
	ticker := time.NewTicker(1 * time.Second)
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
				log.Printf("internal/producer/producer.go: could not write message: %v", err)
			} else {
				log.Printf("internal/producer/producer.go: message sent: %s", messageValue)
			}
		}
	}
}
