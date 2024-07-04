package storage

import (
	"database/sql"
	"fmt"
	"logcollector/internal/config"

	"log/slog"
)

type Postgres struct {
	db *sql.DB
}

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

func (p *Postgres) Close() error {
	return p.db.Close()
}

func (p *Postgres) DB() *sql.DB {
	return p.db
}

func (p *Postgres) Ping() error {
	return p.db.Ping()
}
