package user

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(r fiber.Router) {
	r.Get("/csearch", GlobalSearch)
	r.Get("/gsearch", CustomSearch)

	r.Get("/recent-interactions", GetRecentInteractions)
	r.Patch("/recent-interactions/:id", AddToRecentInteractions)
	r.Delete("/recent-interactions/:id", DeleteFromRecentInteractions)
}

/*
Store this data inside android msgstore.db

user_id | listner_id | last_interacte_at | is_active
*/
