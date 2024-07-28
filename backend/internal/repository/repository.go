package repository

import (
	"logcollector/internal/schemas"
	"logcollector/internal/storage/clickhouse"
)

// Logs представляет интерфейс для работы с логами.
type Logs interface {
	AddLog(log schemas.Logs) error
	GetLogs() ([]schemas.Logs, error)
}

// Repository представляет собой структуру, объединяющую различные интерфейсы репозиториев.
type Repository struct {
	Logs
}

// NewRepository создает новый экземпляр Repository с предоставленной базой данных ClickHouse.
func NewRepository(db *clickhouse.ClickHouse) *Repository {
	return &Repository{
		Logs: NewLogsRepository(db),
	}
}
