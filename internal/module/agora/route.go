package agora

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(r fiber.Router) {
	r.Post("/token", GenerateToken)
}
