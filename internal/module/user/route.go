package user

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router) {
	router.Get("/csearch", GlobalSearch)
	router.Get("/gsearch", CustomSearch)

	router.Get("/recent-interactions", GetRecentInteractions)
	router.Patch("/recent-interactions/:id", AddToRecentInteractions)
	router.Delete("/recent-interactions/:id", DeleteFromRecentInteractions)
}

/*
Store this data inside android msgstore.db

user_id | listner_id | last_interacte_at | is_active
*/
