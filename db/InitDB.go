package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB() *sql.DB {
	connStr := "user=postgres password=123 dbname=v1_db_helpers sslmode=disable"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	CreateTables(db)

	return db
}
