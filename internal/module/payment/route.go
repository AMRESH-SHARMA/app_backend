// package payment

// import "github.com/gofiber/fiber/v2"

// func RegisterRoutes(r fiber.Router) {
// 	r.Get("/recharge/options", GetRechargeOptions)
// 	r.Get("/balance/:id", GetBalance)
// 	r.Post("/balance/add/:id", AddBalance)
// 	r.Get("/history/:id", GetHistory)
// }


//------ above hint is provided by amresh



package payment

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(r fiber.Router, c *PaymentController) {

	r.Get("/recharge/options", c.GetRechargeOptions)
	r.Get("/balance/:id", c.GetBalance)
	r.Get("/history/:id", c.GetHistory)

	r.Post("/order", c.CreateOrder)
	r.Post("/intent", c.CreatePaymentIntent)

	r.Post("/webhook", c.Webhook)
	r.Post("/refund", c.CreateRefund)

}

