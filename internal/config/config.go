package config

import (
	"log/slog"
	"os"

	"gopkg.in/yaml.v2"
)

// Config содержит конфигурацию для всех используемых сервисов
type Config struct {
	Kafka    KafkaConfig    `yaml:"kafka"`
	Redis    RedisConfig    `yaml:"redis"`
	Postgres PostgresConfig `yaml:"postgres"`
	API      APIConfig      `yaml:"api"`
}

// KafkaConfig содержит конфигурацию для Kafka
type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
	Group   string   `yaml:"group"`
}

// RedisConfig содержит конфигурацию для Redis
type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// PostgresConfig содержит конфигурацию для PostgreSQL
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

// APIConfig содержит конфигурацию для API
type APIConfig struct {
	Port int `yaml:"port"`
}

// LoadConfig загружает конфигурацию из файла config.yaml
func LoadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		slog.Error("internal/config/config.go: Couldn't read config.yaml: ", "error", err)
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		slog.Error("internal/config/config.go: Couldn't unmarshal config.yaml: ", "error", err)
		return nil, err
	}

	return &cfg, nil
}
