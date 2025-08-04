package database

import (
	"fmt"
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Connect(dsn string, migrationsPath string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Database connection established.")

	if err := runMigrations(db, migrationsPath); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func runMigrations(db *sql.DB, migrationsPath string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return fmt.Errorf("could not create postgres driver for migrations: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Database migrations applied successfully.")
	return nil
}
