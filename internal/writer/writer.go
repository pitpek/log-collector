package producer

import (
	"context"
	"log"
	"logcollector/internal/config"

	kafkaProducer "logcollector/pkg/kafka"

	"github.com/segmentio/kafka-go"
)

// Writer представляет собой структуру для отправки сообщений в Kafka.
type Writer struct {
	writer *kafka.Writer
	key    string
}

// NewWriter создает новый экземпляр Writer с предоставленной конфигурацией Kafka.
// cfg - конфигурация Kafka, содержащая список брокеров, тему и ключ.
func NewWriter(cfg *config.KafkaConfig) *Writer {
	return &Writer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(cfg.Brokers...),
			Topic:    cfg.Topic,
			Balancer: &kafka.LeastBytes{},
		},
		key: cfg.Key,
	}
}

// Start запускает процесс отправки сообщений в Kafka.
// ctx - контекст для управления жизненным циклом процесса.
func (w *Writer) Start(ctx context.Context) {
	prod := kafkaProducer.NewProducer(w.writer)
	go func() {
		if err := prod.Start(ctx, w.key); err != nil {
			log.Fatalf("internal/writer/writer.go: Failed to start writer: %v", err)
		}
	}()
}
