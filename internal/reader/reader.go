package consumer

import (
	"context"
	"log"

	"logcollector/internal/config"
	"logcollector/internal/storage/clickhouse"
	kafkaConsumer "logcollector/pkg/kafka"

	"github.com/segmentio/kafka-go"
)

type Reader struct {
	reader *kafka.Reader
	db     *clickhouse.ClickHouse
}

func NewReader(cfg *config.KafkaConfig, db *clickhouse.ClickHouse) *Reader {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		Topic:   cfg.Topic,
	})

	return &Reader{
		reader: kafkaReader,
		db:     db,
	}
}

func (r *Reader) Start(ctx context.Context) {
	cons := kafkaConsumer.NewConsumer(r.reader, r.db)
	go func() {
		if err := cons.Start(ctx); err != nil {
			log.Fatalf("internal/consumer/consumer.go: Failed to start consumer: %v", err)
		}
	}()
}

func (r *Reader) Stop() {
	if err := r.reader.Close(); err != nil {
		log.Printf("internal/consumer/consumer.go: Failed to close reader: %v", err)
	}
}
