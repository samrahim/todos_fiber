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
	app.Get("/todos/:id", routes.GetTodoById)
	app.Delete("/todos/:id", routes.DeleteTodo)
	app.Put("/todos/:id", routes.UpdateTodo)
}

func main() {
	database.ConnectToDb()
	app := fiber.New()
	SetUpRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
