package helpers

import (
	"database/sql"
	"log"
	gstructures "v1/test/fiber/db/helpers/GStructures"
)

func HelperEditUser(db *sql.DB, EditableUser_Id int, EditFields gstructures.GEditUser) map[string]string {
	query := "UPDATE users SET username = $1 WHERE id = $2"

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	_, err = statement.Exec(EditFields.Username, EditableUser_Id)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	response := map[string]string{"status": "success"}
	return response
}
