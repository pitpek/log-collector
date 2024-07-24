package writer

import (
	"context"
	"log"
	"logcollector/internal/config"

	kafkaProducer "logcollector/pkg/kafka"

	"github.com/segmentio/kafka-go"
)

// Writer представляет собой структуру для отправки сообщений в Kafka.
type Writer struct {
	kafkaProducer kafkaProducer.Producer
}

// NewWriter создает новый экземпляр Writer с предоставленной конфигурацией Kafka.
// cfg - конфигурация Kafka, содержащая список брокеров, тему и ключ.
func NewWriter(cfg *config.KafkaConfig) *Writer {
	producer := kafkaProducer.NewProducer(&kafka.Writer{
		Addr:     kafka.TCP(cfg.Brokers...),
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &Writer{
		kafkaProducer: *producer,
	}
}

// Start запускает процесс отправки сообщений в Kafka.
// ctx - контекст для управления жизненным циклом процесса.
func (w *Writer) Start(ctx context.Context, key string) {
	go func() {
		if err := w.kafkaProducer.Start(ctx, key); err != nil {
			log.Fatalf("internal/witer/writer.go: Failed to start writer: %v", err)
		}
	}()
}
