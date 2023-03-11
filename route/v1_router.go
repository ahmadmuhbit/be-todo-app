package route

import (
	"be-todo-app/controller"

	"github.com/gofiber/fiber/v2"
)

func v1Router(app *fiber.App) {
	v1 := app.Group("/api/v1")

	todo := v1.Group("/todos")

	todo.Post("/", controller.Create)
	todo.Get("/", controller.GetAll)
	todo.Get("/:id", controller.GetById)
	todo.Patch("/:id", controller.Update)
	todo.Delete("/:id", controller.Delete)
}
