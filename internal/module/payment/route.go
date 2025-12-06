package payment

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(r fiber.Router) {
	r.Get("/recharge/options", GetRechargeOptions)
	r.Get("/balance/:id", GetBalance)
	r.Post("/balance/add/:id", AddBalance)
	r.Get("/history/:id", GetHistory)
}
