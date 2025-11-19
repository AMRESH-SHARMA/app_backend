package internal

import (
	"app_backend/internal/module/listener"
	"app_backend/internal/module/test"

	// import other modules as needed
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func StartServer() error {
	app := fiber.New()
	test.RegisterRoutes(app)
	// Register module routes under /api/v1/...
	api := app.Group("/api/v1")
	listener.RegisterRoutes(api.Group("/listeners"))
	// user.RegisterRoutes(api.Group("/users"))

	port := viper.GetInt("port")
	if port == 0 {
		port = 8080
	}
	return app.Listen(fmt.Sprintf(":%d", port))
}
