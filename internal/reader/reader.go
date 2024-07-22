package reader

import (
	"context"
	"log"
	"logcollector/internal/config"
	"logcollector/internal/repository"
	kafkaConsumer "logcollector/pkg/kafka"

	"github.com/segmentio/kafka-go"
)

// Reader представляет собой структуру для чтения сообщений из Kafka и сохранения их в базу данных.
type Reader struct {
	kafkaConsumer *kafkaConsumer.Consumer
}

// NewReader создает новый экземпляр Reader с предоставленной конфигурацией Kafka и репозиторием базы данных.
func NewReader(cfg *config.KafkaConfig, db *repository.Repository) *Reader {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		Topic:   cfg.Topic,
	})

	consumer := kafkaConsumer.NewConsumer(kafkaReader, db)
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

// Stop останавливает чтение сообщений из Kafka и закрывает reader.
func (r *Reader) Stop() {
	if err := r.kafkaConsumer.Stop(); err != nil {
		log.Printf("internal/reader/reader.go: Failed to close reader: %v", err)
	}
}
