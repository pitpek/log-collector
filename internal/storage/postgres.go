package storage

import (
	"database/sql"
	"fmt"
	"logcollector/internal/config"

	"log/slog"
)

// Postgres представляет собой структуру, которая оборачивает соединение с базой данных PostgreSQL
type Postgres struct {
	db *sql.DB
}

// NewPostgres создает новое соединение с базой данных PostgreSQL
// cfg - конфигурация подключения к базе данных PostgreSQL
func NewPostgres(cfg *config.PostgresConfig) (*Postgres, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slog.Error("internal/storage/postgres.go: Failed to start postgres: ", "error", err)
		return nil, err
	}
	return &Postgres{db: db}, nil
}

// Close закрывает соединение с базой данных
func (p *Postgres) Close() error {
	return p.db.Close()
}

// DB возвращает экземпляр базы данных SQL
func (p *Postgres) DB() *sql.DB {
	return p.db
}

// Ping проверяет соединение с базой данных
func (p *Postgres) Ping() error {
	return p.db.Ping()
}

// InsertMessage вставляет сообщение в таблицу logs
// message - сообщение, которое нужно вставить
func (p *Postgres) InsertMessage(message string) error {
	_, err := p.DB().Exec(`INSERT INTO logs (message) VALUES ($1)`, message)
	return err
}
