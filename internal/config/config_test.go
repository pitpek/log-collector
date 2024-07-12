package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Создаем временный файл конфигурации
	tempFile, err := os.CreateTemp("", "config.yaml")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	// Записываем пример конфигурации в файл
	_, err = tempFile.Write([]byte(`
api:
  port: 8080
kafka:
  brokers: ["kafka:9092"]
  topic: log-messages
  key: log_collector
redis:
  address: redis:6379
  password: ""
  db: 0
clickhouse:
  host: clickhouse
  port: 9000
  user: clickhouse
  password: clickhouse
  dbname: log_collector
`))
	assert.NoError(t, err)
	tempFile.Close()

	// Загружаем конфигурацию
	cfg, err := LoadConfig(tempFile.Name())
	assert.NoError(t, err)

	// Проверяем загруженные данные
	assert.Equal(t, 8080, cfg.API.Port)
	assert.Equal(t, "kafka:9092", cfg.Kafka.Brokers[0])
	assert.Equal(t, "log-messages", cfg.Kafka.Topic)
	assert.Equal(t, "log_collector", cfg.Kafka.Key)
	assert.Equal(t, "redis:6379", cfg.Redis.Address)
	assert.Equal(t, "", cfg.Redis.Password)
	assert.Equal(t, 0, cfg.Redis.DB)
	assert.Equal(t, "clickhouse", cfg.ClickHouse.Host)
	assert.Equal(t, 9000, cfg.ClickHouse.Port)
	assert.Equal(t, "clickhouse", cfg.ClickHouse.User)
	assert.Equal(t, "clickhouse", cfg.ClickHouse.Password)
	assert.Equal(t, "log_collector", cfg.ClickHouse.DBName)
}
