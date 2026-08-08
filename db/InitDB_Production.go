package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB_Production() *sql.DB {
	connStr := "user= password= dbname= sslmode="
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	CreateTables(db)

	return db
}
