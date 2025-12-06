package user

import (
	"app_backend/internal/database"
	"app_backend/internal/response"

	"github.com/gofiber/fiber/v2"
)

func UpdateDeviceToken(c *fiber.Ctx) error {
	req := new(DeviceTokenRequest)

	// Parse JSON
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request body", fiber.StatusBadRequest)
	}

	if req.UserID == "" || req.DeviceToken == "" {
		return response.Error(c, "UserId and DeviceToken required", fiber.StatusBadRequest)
	}

	// Update DB
	if err := database.DB.
		Model(&User{}).
		Where("id = ?", req.UserID).
		Update("device_token", req.DeviceToken).
		Error; err != nil {
		return response.Error(c, "Failed to update device token", fiber.StatusInternalServerError)
	}

	return response.Success(c, fiber.Map{
		"userId":      req.UserID,
		"deviceToken": req.DeviceToken,
	}, "Device token updated", fiber.StatusOK)
}
