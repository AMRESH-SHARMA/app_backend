package internal

import (
	"app_backend/internal/database"
	"app_backend/internal/module/call"
	"app_backend/internal/module/listener"
	"app_backend/internal/module/notification"
	"app_backend/internal/module/payment"
	"app_backend/internal/module/rtc"
	"app_backend/internal/module/test"
	"app_backend/internal/module/user"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"app_backend/internal/seed"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func StartServer() error {

	database.ConnectDB()

	if viper.GetString("ENV") == "development" {
		/*
			Uncomment to enable DB Migration and Seeding
			After uncomment need to import using quick fix
		*/
		database.DB.AutoMigrate(&user.User{})
		database.DB.AutoMigrate(&listener.Listener{})

		seed.Run(database.DB)
	}

	app := fiber.New()

	if viper.GetString("ENV") == "development" {
		app.Use(logger.New(logger.Config{
			Format:     "${time} | ${ip} | ${status} | ${latency} | ${method} ${path}\n",
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Local",
		}))

		app.Use(func(c *fiber.Ctx) error {
			fmt.Printf("\nREQ BODY: %s\n", c.Body())
			return c.Next()
		})
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	notification.InitFCM()

	// app := fiber.New(fiber.Config{
	// 	ErrorHandler: response.Error,
	// })

	test.RegisterRoutes(app)
	// Register module routes under /api/v1/...
	v1 := app.Group("/api/v1")
	listener.RegisterRoutes(v1.Group("/listeners"))
	payment.RegisterRoutes(v1.Group("/payments"))
	user.RegisterDeviceTokenRoutes(v1.Group("/user"))

	// RTC [Agora]
	rtc.RegisterRoutes(v1.Group("/rtc"))
	call.RegisterRoutes(v1.Group("/call"))

	port := viper.GetInt("PORT")
	if port == 0 {
		port = 8080
	}
	// return app.Listen(fmt.Sprintf(":%d", port))
	return app.Listen(fmt.Sprintf("0.0.0.0:%d", port))
}
