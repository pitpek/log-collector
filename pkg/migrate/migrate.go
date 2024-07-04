package migrate

import (
	"database/sql"
	"log"

	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func StartMigration(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.Error("Could not create database driver: ", "error", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://scripts/migrations",
		"postgres", driver)
	if err != nil {
		slog.Error("Could not create migrate instance: ", "error", err)
		return err
	}

	// Apply the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error("Could not run up migrations: ", "error", err)
		return err
	}

	log.Println("Migrations applied successfully")
	return nil
}
