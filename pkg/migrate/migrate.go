package migrate

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// импортируем драйвер файловой системы для миграций
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// StartMigration выполняет миграции базы данных PostgreSQL
// db - открытое соединение с базой данных PostgreSQL
func StartMigration(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Printf("pkg/migrate/migrate.go: Could not create database driver: %v", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://scripts/migrations",
		"postgres", driver)
	if err != nil {
		log.Printf("pkg/migrate/migrate.go: Could not create migrate instance: %v", err)
		return err
	}

	// Применение миграций
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("pkg/migrate/migrate.go: Could not run up migrations: %v", err)
		return err
	}

	log.Println("pkg/migrate/migrate.go: Migrations applied successfully")
	return nil
}
