package db

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			ID SERIAL PRIMARY KEY,
			UUID TEXT NOT NULL,
			USERNAME VARCHAR(100) NOT NULL,
			PASSWORD TEXT NOT NULL
		)`,
	}

	for _, query := range queries {
		statement, err := db.Prepare(query)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		_, err = statement.Exec()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
