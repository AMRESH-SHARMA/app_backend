package user

import "github.com/gofiber/fiber/v2"

func RegisterDeviceTokenRoutes(r fiber.Router) {
	// Register / Update
	r.Post("/device-token", UpdateDeviceToken)

	// Delete refresh not required
	// r.Delete("/device-token", DeleteDeviceToken)
	// r.Post("/device-token/refresh", RefreshDeviceToken)
}
