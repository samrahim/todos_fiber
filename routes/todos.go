package routes

import (
	"errors"
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
func GetAllTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.Database.DB.Find(&todos)
	todoss := []Todo{}
	for _, todo := range todos {
		response := CreateTodoResponse(todo)
		todoss = append(todoss, response)
	}
	return c.Status(200).JSON(todoss)
}

func find(id int, todo *models.Todo) error {
	database.Database.DB.Find(&todo, "id=?", id)
	if todo.Id == 0 {
		return errors.New("please check your id")
	}
	return nil
}
func GetTodoById(c *fiber.Ctx) error {
	var todo models.Todo
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := find(id, &todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	repsonse := CreateTodoResponse(todo)
	return c.Status(200).JSON(repsonse)
}

func DeleteTodo(c *fiber.Ctx) error {
	var todo models.Todo
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := find(id, &todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.DB.Delete(&todo)
	return c.Status(200).JSON("successefully deleted")
}

func UpdateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := find(id, &todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type Up struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
	}
	var up Up
	if err := c.BodyParser(&up); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	todo.Subtitle = up.Subtitle
	todo.Title = up.Title
	response := CreateTodoResponse(todo)
	return c.Status(200).JSON(response)
}
