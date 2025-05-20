package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // ✅ import driver PostgreSQL
)

const (
	driver   = "postgres"
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "money_management"
)

func InitDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// RunMigrations creates necessary tables if they don't exist
func RunMigrations(db *sql.DB) error {
	// Create transactions table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			type VARCHAR(10) NOT NULL CHECK (type IN ('income', 'expense')),
			amount DECIMAL(12, 2) NOT NULL,
			category VARCHAR(100) NOT NULL,
			description TEXT,
			date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}
