package main

import (
	"log"
	"logcollector/internal/config"
	"logcollector/internal/storage"
	"logcollector/pkg/migrate"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cmd/consumer/main.go: Failed to load config: %v", err)
	}

	db, err := storage.NewPostgres(&cfg.Postgres)
	if err != nil {
		log.Fatalf("cmd/main.go: Failed to start postgres: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("cmd/main.go: Failed to ping postgres: %v", err)
	}
	log.Println("cmd/main.go: Database connected")

	err = migrate.StartMigration(db.DB())
	if err != nil {
		log.Fatalf("cmd/main.go: Failed to run migrations: %v", err)
	}
	log.Println("cmd/main.go: Migrations applied successfully")
}

