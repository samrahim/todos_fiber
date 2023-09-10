package main

import (
	"log"
	"todos/database"
	"todos/routes"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/todos", routes.CreateTodo)
	app.Get("/todos", routes.GetAllTodos)
}

func main() {
	database.ConnextToDb()
	app := fiber.New()
	SetUpRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
