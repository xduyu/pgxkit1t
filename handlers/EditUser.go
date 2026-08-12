package handlers

import (
	"database/sql"
	"v1/test/fiber/db/helpers/db/helpers"
	structures "v1/test/fiber/db/helpers/handlers/Structures"

	"github.com/gofiber/fiber/v3"
)

func EditUser(database *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req structures.HReqEditUser
		c.Bind().Body(&req)
		resp := helpers.HelperEditUser(database, req.Editable_UserId, req.Data)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"info": resp,
		})
	}
}
