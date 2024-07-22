package main

import (
	"context"
	"log"
	"logcollector/internal/api"
	"logcollector/internal/config"
	"logcollector/internal/reader"
	"logcollector/internal/repository"
	"logcollector/internal/server"
	"logcollector/internal/service"
	"logcollector/internal/storage/clickhouse"
	"logcollector/internal/storage/redis"
	"logcollector/internal/writer"
	"logcollector/pkg/migrate"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Инициализация всех зависимостей и возврат ошибки, если что-то пошло не так
func initialize(cfg *config.Config) (*clickhouse.ClickHouse, *redis.Client, *reader.Reader, *writer.Writer, *api.Router, *server.Server, error) {
	db, err := clickhouse.NewClickHouse(&cfg.ClickHouse)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	err = migrate.StartMigration(db.DB())
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	redisClient, err := redis.NewClient(&cfg.Redis)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	err = redisClient.Ping()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := api.NewRouter(service)

	reader := reader.NewReader(&cfg.Kafka, repo)
	writer := writer.NewWriter(&cfg.Kafka)

	serv := new(server.Server)

	return db, redisClient, reader, writer, handler, serv, nil
}

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to load config: %v", err)
	}

	db, redisClient, reader, writer, handler, serv, err := initialize(cfg)

	if err != nil {
		log.Fatalf("cmd/app/main.go: Failed to initialize dependencies: %v", err)
	}
	defer db.Close()
	defer redisClient.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reader.Start(ctx)

	writer.Start(ctx, cfg.Kafka.Key)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("tick tick tick tick tick tick")
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		if err := serv.Run(cfg.API.Port, handler.InitRoutes()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("cmd/app/main.go: Failed to start server: %v", err)
		}
	}()
	log.Println("cmd/app/main.go: Server started successfully")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	log.Println("cmd/app/main.go: Shutting down app")
	cancel()
}
