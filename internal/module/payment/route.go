package payment

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router) {
	router.Get("/recharge/options", GetRechargeOptions)
	router.Get("/balance/:id", GetBalance)
	router.Post("/balance/add/:id", AddBalance)
	router.Get("/history/:id", GetHistory)
}
