package main

import (
	"context"
	"log"
	"logcollector/internal/config"
	"logcollector/internal/consumer"
	"logcollector/internal/storage"
	"logcollector/pkg/migrate"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to load config: %v", err)
	}

	// Инициализация PostgreSQL
	db, err := storage.NewPostgres(&cfg.Postgres)
	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to start postgres: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to ping postgres: %v", err)
	}
	log.Println("cmd/app/main.go: Database connected")

	err = migrate.StartMigration(db.DB())
	if err != nil {
		log.Fatalf("cmd/main.go: Failed to run migrations: %v", err)
	}
	log.Println("cmd/main.go: Migrations applied successfully")

	cons := consumer.NewConsumer(cfg.Kafka.Brokers, cfg.Kafka.Topic, cfg.Kafka.Group, db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := cons.Start(ctx); err != nil {
			log.Fatalf("cmd/app/start.go: Failed to start consumer: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	log.Println("shutting down app")
	cancel()
}
