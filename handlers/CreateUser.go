package handlers

import (
	"database/sql"
	"log"
	helpers "v1/test/fiber/db/helpers/db/helpers"
	checkers "v1/test/fiber/db/helpers/handlers/Checkers"
	structures "v1/test/fiber/db/helpers/handlers/Structures"

	"github.com/gofiber/fiber/v3"
)

func CreateUser(database *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var u structures.HCreateUser
		if err := c.Bind().Body(&u); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err})
		}

		hashedPassword, err := checkers.HashPassword(u.Password)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		helpers.CreateUser(database, u.Username, hashedPassword)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Пользователь успешно создан",
		})
	}
}
