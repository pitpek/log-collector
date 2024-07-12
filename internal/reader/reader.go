package consumer

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
	reader *kafka.Reader
	db     *repository.Repository
}

// NewReader создает новый экземпляр Reader с предоставленной конфигурацией Kafka и репозиторием базы данных.
func NewReader(cfg *config.KafkaConfig, db *repository.Repository) *Reader {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		Topic:   cfg.Topic,
	})

	return &Reader{
		reader: kafkaReader,
		db:     db,
	}
}

// Start запускает чтение сообщений из Kafka и их обработку.
func (r *Reader) Start(ctx context.Context) {
	cons := kafkaConsumer.NewConsumer(r.reader, r.db)
	go func() {
		if err := cons.Start(ctx); err != nil {
			log.Fatalf("internal/reader/reader.go: Failed to start reader: %v", err)
		}
	}()
}

// Stop останавливает чтение сообщений из Kafka и закрывает reader.
func (r *Reader) Stop() {
	if err := r.reader.Close(); err != nil {
		log.Printf("internal/reader/reader.go: Failed to close reader: %v", err)
	}
}
