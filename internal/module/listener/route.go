package listener

import "github.com/gofiber/fiber/v2"

// RegisterRoutes sets up listener-related routes
func RegisterRoutes(router fiber.Router) {
	router.Post("/register", RegisterListener)
	router.Post("/login", LoginListener)
	router.Get("/profile", GetProfile)
	router.Put("/profile", UpdateProfile)
	router.Get("/availability", GetAvailability)
	router.Put("/availability", UpdateAvailability)
	router.Get("/sessions", GetSessions)
	router.Get("/earnings", GetEarnings)
	router.Get("/ratings", GetRatings)
	router.Get("/history", GetHistory)
	router.Post("/withdraw", WithdrawEarnings)
}
