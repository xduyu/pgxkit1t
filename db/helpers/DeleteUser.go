package helpers

import (
	"database/sql"
	"log"
)

func HelperDeleteUser(db *sql.DB, User_Id int) map[string]string {
	query := "DELETE FROM users WHERE id = $1"

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	_, err = statement.Exec(User_Id)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	response := map[string]string{"status": "success"}
	return response
}
func HelperDeleteUserByUuid(db *sql.DB, User_Uuid string) map[string]string {
	query := "DELETE FROM users WHERE uuid = $1"

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	_, err = statement.Exec(User_Uuid)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	response := map[string]string{"status": "success"}
	return response
}
