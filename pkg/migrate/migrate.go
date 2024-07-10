package migrate

import (
	"database/sql"
	"log"

	// импортируем драйвер файловой системы для миграций
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// StartMigration создает необходимые таблицы в базе данных
func StartMigration(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			date DateTime,
			message String
		) ENGINE = MergeTree()
		ORDER BY date
	`)
	if err != nil {
		log.Printf("internal/storage/clickhouse.go: Failed to create table logs in ClickHouse: %v", err)
	}
	return err
}
