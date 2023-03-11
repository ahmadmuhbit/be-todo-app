package controller

import (
	"be-todo-app/database"
	"be-todo-app/model"
	"be-todo-app/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	todoRequest := request.TodoCreateRequest{}

	// Parse request body
	if errParse := c.BodyParser(&todoRequest); errParse != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "GAGAL PARSING DATA",
			"error":   errParse.Error(),
		})
	}

	// Validation Request Data
	validate := validator.New()
	if errValidate := validate.Struct(&todoRequest); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   errValidate.Error(),
		})
	}

	todo := model.Todo{}
	todo.Name = todoRequest.Name
	todo.IsComplete = todoRequest.IsComplete
	if todoRequest.Note != "" {
		todo.Note = &todoRequest.Note
	}

	if errDb := database.DB.Create(&todo).Error; errDb != nil {
		log.Println("todo.controller.go => Create :: ", errDb)
		return c.Status(500).JSON(fiber.Map{
			"message": "INTERNAL SERVER ERROR",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "OK",
		"data":    todo,
	})
}

func GetAll(c *fiber.Ctx) error {
	todos := []model.Todo{}
	if err := database.DB.Find(&todos).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "INTERNAL SERVER ERROR",
		})
	}
	return c.JSON(fiber.Map{
		"message": "OK",
		"data":    todos,
	})
}

func GetById(c *fiber.Ctx) error {
	todoId := c.Params("id")
	todo := model.Todo{}

	if err := database.DB.First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "NOT FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"message": "OK",
		"data":    todo,
	})
}

func Update(c *fiber.Ctx) error {
	todoRequest := request.TodoUpdateRequest{}

	// Parse request body
	if errParse := c.BodyParser(&todoRequest); errParse != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "GAGAL PARSING DATA",
			"error":   errParse.Error(),
		})
	}

	// Validation Request Data
	validate := validator.New()
	if errValidate := validate.Struct(&todoRequest); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "BAD REQUEST",
			"error":   errValidate.Error(),
		})
	}

	todoId := c.Params("id")
	todo := model.Todo{}

	if err := database.DB.First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "NOT FOUND",
		})
	}

	todo.Name = todoRequest.Name
	todo.Note = &todoRequest.Note
	todo.IsComplete = todoRequest.IsComplete

	if errUpdate := database.DB.Save(&todo).Error; errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "INTERNAL SERVER ERROR",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "OK",
		"data":    todo,
	})
}

func Delete(c *fiber.Ctx) error {
	todoId := c.Params("id")
	todo := model.Todo{}

	if err := database.DB.First(&todo, "id = ?", todoId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "NOT FOUND",
		})
	}

	if errDelete := database.DB.Delete(&todo).Error; errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "INTERNAL SERVER ERROR",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "OK",
	})
}
