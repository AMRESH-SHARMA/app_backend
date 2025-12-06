package listener

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(r fiber.Router) {
	r.Get("/gsearch", GlobalSearch)
	// r.Get("/csearch", CustomSearch)

	r.Post("/register", RegisterListener)
	r.Post("/login", LoginListener)

	r.Get("/", GetAllListener)
	r.Get("/profile/:id", GetProfile)
	r.Put("/profile/:id", UpdateProfile)

	r.Get("/availability/:id", GetAvailability)
	r.Put("/availability/:id", UpdateAvailability)

	// r.Get("/sessions/:id", GetSessions)
	// r.Get("/earnings/:id", GetEarnings)

	r.Get("/ratings/:id", GetRatings)
	r.Get("/history/:id", GetHistory)
	r.Post("/withdraw/:id", WithdrawEarnings)
}
