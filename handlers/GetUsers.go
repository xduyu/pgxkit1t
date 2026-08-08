package handlers

import (
	"database/sql"
	helpers "v1/test/fiber/db/helpers/db/helpers"

	"github.com/gofiber/fiber/v3"
)

func GetUsers(database *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		if helpers.GetUsers(database) == nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"Status": "Not Found",
				"Code":   404,
				"Users":  helpers.GetUsers(database),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Status": "Found Successful",
			"Code":   200,
			"Users":  helpers.GetUsers(database),
		})
	}
}
