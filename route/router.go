package route

import "github.com/gofiber/fiber/v2"

func NewRouter(app *fiber.App) {
	v1Router(app)
}
