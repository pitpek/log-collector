package main

import (
	"context"
	"log"
	"logcollector/internal/api"
	"logcollector/internal/config"
	"logcollector/internal/consumer"
	"logcollector/internal/server"
	"logcollector/internal/storage/clickhouse"
	"logcollector/internal/storage/redis"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Загрузка конфигов
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to load config: %v", err)
	}

	// Инициализация ClickHouse
	db, err := clickhouse.NewClickHouse(&cfg.ClickHouse)
	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to start clickhouse: %v", err)
	}
	defer db.Close()

	// Проверка подключения к ClickHouse
	err = db.Ping()
	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to ping clickhouse: %v", err)
	}
	log.Println("cmd/app/main.go: Database clickhouse connected")

	// Создание таблиц в Clickhouse
	err = clickhouse.CreateTables(db.DB())
	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to create clickhouse tables: %v", err)
	}
	log.Println("cmd/app/main.go: Tables clickhouse created successfully")

	// Инициализация Redis
	redisClient, err := redis.NewClient(&cfg.Redis)
	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to connect redis: %v", err)
	}
	defer redisClient.Close()

	// Проверка подключения к redis
	err = redisClient.Ping()
	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to ping Redis: %v", err)
	}
	log.Println("cmd/app/main.go: Redis connected successfully")

	// Инициализация Consumer
	cons := consumer.NewConsumer(cfg.Kafka.Brokers, cfg.Kafka.Topic, cfg.Kafka.Group, db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := cons.Start(ctx); err != nil {
			log.Fatalf("cmd/app/main.go: Failed to start consumer: %v", err)
		}
	}()
	log.Println("cmd/app/main.go: Consumer started successfully")

	// Запуск HTTP-сервера
	serv := new(server.Server)
	go func() {
		if err := serv.Run(cfg.API.Port, api.InitRoutes()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("cmd/app/main.go: Server start error: %v", err)
		}
	}()
	log.Println("cmd/app/main.go: Server started successfully")

	// Ожидание комманды завершения программы
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	log.Println("cmd/app/main.go: Shutting down app")
	cancel()
}
