package main

import (
	"v1/test/fiber/db/helpers/db"
	"v1/test/fiber/db/helpers/handlers"
	"v1/test/fiber/db/helpers/middlewares"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	database := db.InitDB()
	defer database.Close()
	api := app.Group("/api/v1", middlewares.TimeLogger)

	// For example we try to use users route
	api.Get("/users", handlers.GetUsers(database))
	api.Post("/user", handlers.CreateUser(database))
	api.Delete("/user/id", handlers.DeleteUserById(database))
	api.Delete("/user/uuid", handlers.DeleteUserByUuid(database))
	api.Patch("/user/id", handlers.EditUser(database))

	app.Listen(":3030")
}
