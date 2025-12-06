package call

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(r fiber.Router) {
	r.Post("/start", StartCall)
	r.Post("/accept", AcceptCall)
	r.Post("/reject", RejectCall)
}
