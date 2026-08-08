package handlers

import (
	"database/sql"
	helpers "v1/test/fiber/db/helpers/db/helpers"
	structures "v1/test/fiber/db/helpers/handlers/Structures"

	"github.com/gofiber/fiber/v3"
)

func DeleteUserById(database *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var u_id structures.HDeleteUser
		if err := c.Bind().Body(&u_id); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"UserId":    "NotFound",
				"ErrorCode": fiber.StatusNotFound,
			})
		}

		resp := helpers.HelperDeleteUser(database, u_id.Id)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"info": resp,
		})
	}
}

func DeleteUserByUuid(database *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var u_uuid structures.HDeleteUserByUuid
		if err := c.Bind().Body(&u_uuid); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"UserId":    "NotFound",
				"ErrorCode": fiber.StatusNotFound,
			})
		}
		resp := helpers.HelperDeleteUserByUuid(database, u_uuid.Uuid)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"info": resp,
		})
	}
}
