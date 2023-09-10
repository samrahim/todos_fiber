package main

import (
	"log"
	"todos/database"
	"todos/routes"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/todos", routes.CreateTodo)
}

func main() {
	database.ConnextToDb()
	app := fiber.New()
	SetUpRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
