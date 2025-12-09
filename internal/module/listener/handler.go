package listener

import (
	"app_backend/internal/response"

	"github.com/gofiber/fiber/v2"
)

// global search for listeners by name, ID, or language
func GlobalSearch(c *fiber.Ctx) error {
	accountID := c.Query("account_id")
	name := c.Query("name")
	gender := c.Query("gender")
	lang := c.Query("lang")

	listeners, err := GlobalSearchR(accountID, name, gender, lang)
	if err != nil {
		return response.Error(c, "Failed to search listeners", fiber.StatusInternalServerError)
	}

	if len(listeners) == 0 {
		return response.Success(c, listeners, "No Listeners found matching criteria", fiber.StatusOK)
	}

	return response.Success(c, listeners, "Listeners found", fiber.StatusOK)
}

// // custom search for listeners by language, gender, and age group
// func CustomSearch(c *fiber.Ctx) error {
// 	language := c.Query("language")
// 	gender := c.Query("gender")
// 	ageGroup := c.Query("age_group")
// 	var listeners []testdb.Listener
// 	// if dbInstance == nil {
// 	// 	return response.Success(c, []testdb.Listener{}, "No DB instance", fiber.StatusOK)
// 	// }
// 	// dbInstance.Where("language = ? AND gender = ? AND age_group = ?", language, gender, ageGroup).Find(&listeners)

// 	return response.Success(c, fiber.Map{
// 		"results":   listeners,
// 		"language":  language,
// 		"gender":    gender,
// 		"age_group": ageGroup,
// 	}, "Custom search successful", fiber.StatusOK)
// }

// RegisterListener handles listener registration
func RegisterListener(c *fiber.Ctx) error {
	// TODO: implement registration logic
	return c.JSON(fiber.Map{"message": "Listener registered"})
}

// LoginListener handles listener login
func LoginListener(c *fiber.Ctx) error {
	// TODO: implement login logic
	return c.JSON(fiber.Map{"message": "Listener logged in"})
}

// Add server side pagination
func GetAllListener(c *fiber.Ctx) error {
	users, _ := GetAllListenerR()

	var resp []ListenerGetResponse

	for _, u := range users {
		resp = append(resp, ListenerGetResponse{
			AccountID:   u.AccountID,
			Name:        u.Name,
			Gender:      u.Gender,
			Languages:   u.Languages,
			Age:         u.Listener.Age,
			Avatar:      u.Listener.Avatar,
			TagLine:     u.Listener.TagLine,
			About:       u.Listener.About,
			PricePerMin: u.Listener.PricePerMin,
			Rating:      u.Listener.Rating,
		})
	}

	return response.Success(c, resp, "All Listeners", fiber.StatusOK)
}

func GetProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	/*
		account_id
		name
		avataar
		tag_line
		posts
		rating
		reviewesCount
		experience
		language
	*/
	return c.JSON(fiber.Map{"profile": id})
}

func UpdateProfile(c *fiber.Ctx) error {
	// TODO: implement profile update
	return c.JSON(fiber.Map{"message": "Profile updated"})
}

func GetAvailability(c *fiber.Ctx) error {
	// TODO: implement availability retrieval
	return c.JSON(fiber.Map{"available": true})
}

func UpdateAvailability(c *fiber.Ctx) error {
	// TODO: implement availability update
	return c.JSON(fiber.Map{"message": "Availability updated"})
}

func GetSessions(c *fiber.Ctx) error {
	// TODO: implement session listing
	return c.JSON(fiber.Map{"sessions": []interface{}{}})
}

func GetEarnings(c *fiber.Ctx) error {
	// TODO: implement earnings summary
	return c.JSON(fiber.Map{"earnings": 0})
}

func GetRatings(c *fiber.Ctx) error {
	// TODO: implement ratings retrieval
	return c.JSON(fiber.Map{"ratings": []interface{}{}})
}

func GetHistory(c *fiber.Ctx) error {
	// TODO: implement history retrieval
	return c.JSON(fiber.Map{"history": []interface{}{}})
}

func WithdrawEarnings(c *fiber.Ctx) error {
	// TODO: implement withdrawal logic
	return c.JSON(fiber.Map{"message": "Withdrawal requested"})
}
