package clickhouse

import (
	"database/sql"
	"fmt"
	"log"
	"logcollector/internal/config"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

// ClickHouse представляет собой структуру, которая оборачивает соединение с базой данных ClickHouse
type ClickHouse struct {
	db *sql.DB
}

// NewClickHouse создает новое соединение с базой данных ClickHouse
// cfg - конфигурация подключения к базе данных ClickHouse
func NewClickHouse(cfg *config.ClickHouseConfig) (*ClickHouse, error) {
	connStr := fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&enable_http_compression=1",
		cfg.Host, cfg.Port, cfg.User, cfg.Password)
	db, err := sql.Open("clickhouse", connStr)
	if err != nil {
		log.Printf("internal/storage/clickhouse.go: Failed to start clickhouse: %v", err)
		return nil, err
	}
	return &ClickHouse{db: db}, nil
}

// NewClickHouseWithDB создает новый экземпляр ClickHouse с заданной базой данных
func NewClickHouseWithDB(db *sql.DB) *ClickHouse {
	return &ClickHouse{db: db}
}

// Close закрывает соединение с базой данных
func (c *ClickHouse) Close() error {
	return c.db.Close()
}

// DB возвращает экземпляр базы данных SQL
func (c *ClickHouse) DB() *sql.DB {
	return c.db
}

// Ping проверяет соединение с базой данных
func (c *ClickHouse) Ping() error {
	return c.db.Ping()
}
