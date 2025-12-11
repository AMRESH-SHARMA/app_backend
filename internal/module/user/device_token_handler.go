package user

import (
	"app_backend/internal/database"
	"app_backend/internal/response"

	"github.com/gofiber/fiber/v2"
)

// ------------------- REGISTER / UPDATE ---------------------

func UpdateDeviceToken(c *fiber.Ctx) error {
	req := new(DeviceTokenRequest)

	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request body", fiber.StatusBadRequest)
	}

	if req.AccountID == 0 || req.DeviceToken == "" {
		return response.Error(c, "AccountID and DeviceToken required", fiber.StatusBadRequest)
	}

	// upsert token
	if err := database.DB.
		Model(&User{}).
		Where("id = ?", req.AccountID).
		Update("device_token", req.DeviceToken).Error; err != nil {
		return response.Error(c, "Failed to update device token", fiber.StatusInternalServerError)
	}

	return response.Success(c, fiber.Map{
		"AccountID":   req.AccountID,
		"deviceToken": req.DeviceToken,
	}, "Device token updated", fiber.StatusOK)
}

// ------------------- REMOVE (LOGOUT) ---------------------

// func DeleteDeviceToken(c *fiber.Ctx) error {
// 	req := new(DeviceTokenRequest)

// 	if err := c.BodyParser(req); err != nil {
// 		return response.Error(c, "Invalid request body", fiber.StatusBadRequest)
// 	}

// 	if req.UserID == "" {
// 		return response.Error(c, "UserId required", fiber.StatusBadRequest)
// 	}

// 	// remove token
// 	if err := database.DB.
// 		Model(&User{}).
// 		Where("id = ?", req.UserID).
// 		Update("device_token", "").Error; err != nil {
// 		return response.Error(c, "Failed to delete device token", fiber.StatusInternalServerError)
// 	}

// 	return response.Success(c, nil, "Device token removed", fiber.StatusOK)
// }

// ------------------- REFRESH (optional) ---------------------

// func RefreshDeviceToken(c *fiber.Ctx) error {
// 	req := new(RefreshDeviceTokenRequest)

// 	if err := c.BodyParser(req); err != nil {
// 		return response.Error(c, "Invalid request body", fiber.StatusBadRequest)
// 	}

// 	if req.UserID == "" || req.NewToken == "" {
// 		return response.Error(c, "UserId and NewToken required", fiber.StatusBadRequest)
// 	}

// 	// just overwrite
// 	if err := database.DB.
// 		Model(&User{}).
// 		Where("id = ?", req.UserID).
// 		Update("device_token", req.NewToken).Error; err != nil {
// 		return response.Error(c, "Failed to refresh device token", fiber.StatusInternalServerError)
// 	}

// 	return response.Success(c, fiber.Map{
// 		"userId":      req.UserID,
// 		"deviceToken": req.NewToken,
// 	}, "Device token refreshed", fiber.StatusOK)
// }
