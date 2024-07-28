package clickhouse_test

import (
	"database/sql"
	"testing"
	"time"

	"logcollector/internal/storage/clickhouse"

	"github.com/stretchr/testify/assert"
)

func setupClickHouseTest(t *testing.T) *sql.DB {
	// Используем Docker-образ ClickHouse с хостом "localhost" и портом "9000"
	connStr := "tcp://localhost:9000?username=default&password=&enable_http_compression=1"
	db, err := sql.Open("clickhouse", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to test ClickHouse: %v", err)
	}

	// Создание таблицы logs, если она не существует
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS logs (
		date DateTime,
		app_name String,
		message String
	) ENGINE = MergeTree()
	ORDER BY date;
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("Failed to create logs table: %v", err)
	}

	return db
}

func TestInsertMessageIntegration(t *testing.T) {
	db := setupClickHouseTest(t)
	defer db.Close()

	ch := clickhouse.NewClickHouseWithDB(db)

	date := time.Now()
	key := "testapp"
	message := "test message"

	err := ch.InsertMessage(date, key, message)
	assert.NoError(t, err, "Expected no error when inserting message into ClickHouse")
}
