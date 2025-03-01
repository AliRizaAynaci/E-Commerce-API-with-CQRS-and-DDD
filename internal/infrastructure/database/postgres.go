package database

import (
	"database/sql"
	"e-commerce/pkg/config"
	"fmt"
	"log"

	_ "modernc.org/sqlite"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// NewPostgresConnection creates a new PostgreSQL connection
func NewPostgresConnection(cfg *config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.PostgresConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	log.Println("Connected to PostgreSQL database")
	return db, nil
}

// Close closes the database connection
func Close(db *sql.DB) {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
			return
		}
		log.Println("Database connection closed")
	}
}
