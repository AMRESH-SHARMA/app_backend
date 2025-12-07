package user

import "github.com/gofiber/fiber/v2"

func RegisterDeviceTokenRoutes(r fiber.Router) {
	// Register / Update
	r.Post("/device-token", UpdateDeviceToken)
	r.Delete("/device-token", DeleteDeviceToken)
	r.Post("/device-token/refresh", RefreshDeviceToken)
}
