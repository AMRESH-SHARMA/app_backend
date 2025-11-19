package internal

import (
	"app_backend/internal/module/test"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func StartServer() error {
	app := fiber.New()
	test.RegisterRoutes(app)
	port := viper.GetInt("port")
	if port == 0 {
		port = 8080
	}
	return app.Listen(fmt.Sprintf(":%d", port))
}
