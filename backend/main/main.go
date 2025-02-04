package main

import (
	"go-get-stuff-done/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.ConnectDB()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
