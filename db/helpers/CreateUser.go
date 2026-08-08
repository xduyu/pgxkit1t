package helpers

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func CreateUser(db *sql.DB, username string, password string) map[string]string {
	query := "INSERT INTO users (uuid, username, password) VALUES ($1, $2, $3)"

	NewUserUuid := uuid.New()

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	_, err = statement.Exec(NewUserUuid, username, password)

	response := map[string]string{"status": "success"}
	return response
}
