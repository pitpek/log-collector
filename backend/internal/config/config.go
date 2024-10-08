package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config содержит конфигурацию для всех используемых сервисов
type Config struct {
	Kafka      KafkaConfig      `yaml:"kafka"`
	Redis      RedisConfig      `yaml:"redis"`
	ClickHouse ClickHouseConfig `yaml:"clickhouse"`
	API        APIConfig        `yaml:"api"`
}

// KafkaConfig содержит конфигурацию для Kafka
type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
	Key     string   `yaml:"key"`
}

// RedisConfig содержит конфигурацию для Redis
type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// ClickHouseConfig содержит конфигурацию для ClickHouse
type ClickHouseConfig struct {
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
		log.Printf("internal/config/config.go: Couldn't read config.yaml: %v", err)
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Printf("internal/config/config.go: Couldn't unmarshal config.yaml: %v", err)
		return nil, err
	}

	return &cfg, nil
}
