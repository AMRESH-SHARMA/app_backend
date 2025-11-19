package route

import (
	"app_backend/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	// Default test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server running successfully!")
	})

	// 👋 Sample Hello World API using handler
	app.Get("/hello", handler.HelloHandler)
}
