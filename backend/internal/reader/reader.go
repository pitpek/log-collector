package reader

import (
	"context"
	"log"
	"logcollector/internal/config"
	"logcollector/internal/service"
	kafkaConsumer "logcollector/pkg/kafka"

	"github.com/segmentio/kafka-go"
)

// Reader представляет собой структуру для чтения сообщений из Kafka и сохранения их в базу данных.
type Reader struct {
	kafkaConsumer *kafkaConsumer.Consumer
}

// NewReader создает новый экземпляр Reader с предоставленной конфигурацией Kafka и репозиторием базы данных.
func NewReader(cfg *config.KafkaConfig, service *service.Service) *Reader {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		Topic:   cfg.Topic,
	})

	consumer := kafkaConsumer.NewConsumer(kafkaReader, service)
	return &Reader{
		kafkaConsumer: consumer,
	}
}

// Start запускает чтение сообщений из Kafka и их обработку.
func (r *Reader) Start(ctx context.Context) {
	go func() {
		if err := r.kafkaConsumer.Start(ctx); err != nil {
			log.Fatalf("internal/reader/reader.go: Failed to start reader: %v", err)
		}
	}()
}
