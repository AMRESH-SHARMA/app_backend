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

	// "app_backend/internal/seed"

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
		// database.DB.AutoMigrate(&user.User{})
		// database.DB.AutoMigrate(&listener.Listener{})

		// seed.Run(database.DB)
	}

	app := fiber.New()

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
	return app.Listen(fmt.Sprintf(":%d", port))
}
