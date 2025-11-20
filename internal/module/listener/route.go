package listener

import "github.com/gofiber/fiber/v2"

// RegisterRoutes sets up listener-related routes
func RegisterRoutes(router fiber.Router) {
	router.Get("/gsearch", GlobalSearch)
	router.Get("/csearch", CustomSearch)

	// router.Post("/register", RegisterListener)
	router.Post("/", GetAllListener)
	router.Post("/login", LoginListener)
	router.Get("/profile/:id", GetProfile)
	router.Put("/profile/:id", UpdateProfile)
	router.Get("/availability/:id", GetAvailability)
	router.Put("/availability/:id", UpdateAvailability)
	// router.Get("/sessions/:id", GetSessions)
	// router.Get("/earnings/:id", GetEarnings)
	router.Get("/ratings/:id", GetRatings)
	router.Get("/history/:id", GetHistory)
	router.Post("/withdraw/:id", WithdrawEarnings)
}
