package repository

import (
	"log"
	"logcollector/internal/schemas"
	"logcollector/internal/storage/clickhouse"
	"time"
)

// LogsRepository представляет собой структуру для работы с логами в ClickHouse.
type LogsRepository struct {
	db *clickhouse.ClickHouse
}

// NewLogsRepository создает новый экземпляр LogsRepository с предоставленной базой данных ClickHouse.
func NewLogsRepository(db *clickhouse.ClickHouse) *LogsRepository {
	return &LogsRepository{db: db}
}

// AddLog вставляет сообщение в таблицу logs.
// date - дата и время сообщения
// key - название приложения, с которого пришло сообщение
// message - сообщение, которое нужно вставить
func (lr *LogsRepository) AddLog(date time.Time, key, message string) error {
	_, err := lr.db.DB().Exec("INSERT INTO logs (date, app_name, message) VALUES (?, ?, ?)",
		date, key, message,
	)
	if err != nil {
		log.Printf("internal/storage/clickhouse.go: Failed to insert message into ClickHouse: %v", err)
		return err
	}
	return nil
}

// GetLogs извлекает все сообщения из таблицы logs.
func (lr *LogsRepository) GetLogs() ([]schemas.Logs, error) {
	rows, err := lr.db.DB().Query("SELECT date, app_name, message FROM logs")
	if err != nil {
		log.Printf("internal/storage/clickhouse.go: Failed to query logs from ClickHouse: %v", err)
		return nil, err
	}
	defer rows.Close()

	var logs []schemas.Logs
	for rows.Next() {
		var logRecord schemas.Logs
		if err := rows.Scan(&logRecord.Date, &logRecord.AppName, &logRecord.Message); err != nil {
			log.Printf("internal/storage/clickhouse.go: Failed to scan log row: %v", err)
			return nil, err
		}
		logs = append(logs, logRecord)
	}

	if err := rows.Err(); err != nil {
		log.Printf("internal/storage/clickhouse.go: Rows iteration error: %v", err)
		return nil, err
	}

	return logs, nil
}
