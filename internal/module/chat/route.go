package chat

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(r fiber.Router) {
	// chat := app.Group("/chat")

	r.Post("/session", CreateSession)
	r.Post("/message", SendMessage)
	r.Get("/messages", GetMessages)
	r.Post("/message/status", UpdateMessageStatus)

}
