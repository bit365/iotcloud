package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dsn string) {
	m, err := migrate.New(
		"file://internal/database/migrations",
		dsn,
	)

	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}
	if err := m.Up(); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	log.Println("Migrations applied successfully.")
}
