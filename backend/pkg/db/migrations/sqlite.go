package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	// Open SQLite connection
	db, err := sql.Open("sqlite3", "./social-network.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Apply migrations
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://backend/pkg/db/migrations/sqlite", // Path to migrations
		"sqlite3", // Database type
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize migrations: %v", err)
	}

	// Run all pending migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, fmt.Errorf("failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
	return db, nil
}
