package repository

import (
	"logcollector/internal/schemas"
	"logcollector/internal/storage/clickhouse"
	"time"
)

type LogsRepository struct {
	db *clickhouse.ClickHouse
}

func NewLogsRepository(db *clickhouse.ClickHouse) *LogsRepository {
	return &LogsRepository{db: db}
}

func (lr *LogsRepository) GetLogs() ([]schemas.Logs, error) {
	log := schemas.Logs{
		Date:    time.Now(),
		AppName: "some name",
		Message: "some message",
	}

	logs := []schemas.Logs{log}
	return logs, nil
}
