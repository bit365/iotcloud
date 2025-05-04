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
		log.Println("Failed to create migration instance:", err)
		return
	}

	if err := m.Up(); err != migrate.ErrNoChange {
		log.Println("Failed to apply migrations:", err)
		return
	}

	log.Println("Migrations applied successfully")
}
