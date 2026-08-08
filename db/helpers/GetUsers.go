package helpers

import (
	"database/sql"
	"log"
	helpersstructures "v1/test/fiber/db/helpers/db/helpers/Helpers_Structures"
)

func GetUsers(db *sql.DB) []helpersstructures.HelperGetUser {
	var users []helpersstructures.HelperGetUser
	query := "SELECT id, uuid, username FROM users"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for rows.Next() {
		var u helpersstructures.HelperGetUser
		err := rows.Scan(&u.ID, &u.Uuid, &u.Username)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		users = append(users, u)
	}

	return users
}
