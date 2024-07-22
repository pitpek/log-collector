package migrate

import (
	"database/sql"
	"log"
)

// StartMigration создает необходимые таблицы в базе данных
func StartMigration(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			date DateTime,
			app_name String,
			message String
		) ENGINE = MergeTree()
		ORDER BY date
	`)
	if err != nil {
		log.Printf("pkg/migrate/migrate.go: Failed to create table logs in ClickHouse: %v", err)
	}
	return err
}
