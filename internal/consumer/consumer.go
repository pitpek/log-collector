package consumer

import (
	"logcollector/internal/config"
	"logcollector/pkg/kafka"
)

type Consumer struct {
	cfg *config.Config
}

func NewConsumer(cfg *config.Config) *Consumer {
	return &Consumer{cfg: cfg}
}

func (c *Consumer) Start() error {
	kafkaConsumer := kafka.NewConsumer(c.cfg.Kafka.Brokers, c.cfg.Kafka.Topic)

	return kafkaConsumer.Consume()
}
