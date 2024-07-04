package main

import (
	"log"
	"logcollector/internal/config"
	"logcollector/internal/consumer"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cmd/consumer/main.go: Failed to load config: %v", err)
	}

	cons := consumer.NewConsumer(cfg)
	if err := cons.Start(); err != nil {
		log.Fatalf("cmd/consumer/start.go: Failed to start consumer: %v", err)
	}

}
