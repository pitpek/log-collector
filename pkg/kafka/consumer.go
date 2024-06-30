package kafka

import (
	"log/slog"

	"github.com/IBM/sarama"
)

type Consumer struct {
	Brokers []string
	Topic   string
}

func NewConsumer(brokers []string, topic string) *Consumer {
	return &Consumer{Brokers: brokers, Topic: topic}
}

func (c *Consumer) Consume() error {
	config := sarama.NewConfig()
	master, err := sarama.NewConsumer(c.Brokers, config)
	if err != nil {
		return err
	}
	defer master.Close()

	consumer, err := master.ConsumePartition(c.Topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}
	defer consumer.Close()

	for message := range consumer.Messages() {
		slog.Info("Message received: ", "Info:", string(message.Value))
	}

	return nil
}
