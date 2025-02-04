package main

import (
	"go-get-stuff-done/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/todo_task", handlers.CreateTask)
	app.Delete("/todo_task/:id", handlers.DeleteTask)
	app.Post("/complete_task/:id", handlers.CompleteTask)
	app.Get("/list_tasks", handlers.ListTasks)
	app.Get("/get_next_task", handlers.GetNextTask)
	app.Get("/productivity_report/:date", handlers.ProductivityReport)
}
