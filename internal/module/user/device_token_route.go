package user

import "github.com/gofiber/fiber/v2"

func RegisterDeviceTokenRoutes(r fiber.Router) {
	r.Post("/device-token", UpdateDeviceToken)
}
