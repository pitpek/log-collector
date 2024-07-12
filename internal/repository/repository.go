package repository

import (
	"logcollector/internal/schemas"
	"logcollector/internal/storage/clickhouse"
)

type Logs interface {
	GetLogs() ([]schemas.Logs, error)
}

type Repository struct {
	Logs
}

func NewRepository(db *clickhouse.ClickHouse) *Repository {
	return &Repository{
		Logs: NewLogsRepository(db),
	}
}
