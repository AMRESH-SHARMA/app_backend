package internal

import (
	"app_backend/internal/database"
	"app_backend/internal/module/listener"
	"app_backend/internal/module/payment"
	"app_backend/internal/module/test"
	"app_backend/internal/seed"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func StartServer() error {

	database.ConnectDB()

	if viper.GetString("ENV") == "development" {
		database.DB.AutoMigrate(&listener.Listener{})
		seed.Run(database.DB)
	}

	app := fiber.New()

	// app := fiber.New(fiber.Config{
	// 	ErrorHandler: response.Error,
	// })

	test.RegisterRoutes(app)
	// Register module routes under /api/v1/...
	v1 := app.Group("/api/v1")
	listener.RegisterRoutes(v1.Group("/listeners"))
	payment.RegisterRoutes(v1.Group("/payments"))
	// user.RegisterRoutes(api.Group("/users"))

	port := viper.GetInt("PORT")
	if port == 0 {
		port = 8080
	}
	return app.Listen(fmt.Sprintf(":%d", port))
}
