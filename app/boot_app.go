package app

import (
	"be-todo-app/config"
	"be-todo-app/route"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func BootApp() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if portEnv := os.Getenv("PORT"); portEnv != "" {
		config.PORT = portEnv
	}

	config.BootDatabase()
	config.ConnectDatabase()
	config.RunMigration()

	app := fiber.New()
	route.NewRouter(app)

	app.Listen(config.PORT)
}
