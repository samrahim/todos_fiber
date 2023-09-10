package routes

import (
	"todos/database"
	"todos/models"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

func CreateTodoResponse(todo models.Todo) Todo {
	return Todo{Id: todo.Id, Title: todo.Title, Subtitle: todo.Subtitle}
}

func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.DB.Create(&todo)
	response := CreateTodoResponse(todo)
	return c.Status(200).JSON(response)
}
