package redis

import (
	"context"
	"log"
	"logcollector/internal/config"

	"github.com/go-redis/redis/v8"
)

// Client представляет собой структуру для хранения клиента Redis.
type Client struct {
	client *redis.Client
}

// NewClient создает и возвращает нового клиента Redis.
func NewClient(redisCfg *config.RedisConfig) (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,  // Адрес Redis сервера
		Password: redisCfg.Password, // Пароль, если он установлен
		DB:       redisCfg.DB,       // Номер базы данных
	})

	return &Client{client: client}, nil
}

// Ping проверяет соединение с Redis.
func (r *Client) Ping() error {
	_, err := r.client.Ping(context.Background()).Result()
	return err
}

// Close закрывает соединение с Redis.
func (r *Client) Close() error {
	err := r.client.Close()
	if err != nil {
		log.Printf("internal/storage/redis/redis.go: error closing redis connection: %v", err)
		return err
	}
	return nil
}
