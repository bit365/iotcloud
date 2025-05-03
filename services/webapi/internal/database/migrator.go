package database

import (
	"log"

	"github.com/golang-migrate/migrate"
)

func RunMigrations(dsn string) {
	m, err := migrate.New(
		"file://migrations",
		dsn,
	)

	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	log.Println("Migrations applied successfully.")
}
