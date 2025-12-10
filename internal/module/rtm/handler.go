package rtm

import (
	"app_backend/internal/response"

	"github.com/gofiber/fiber/v2"
)

func SendRTMMessage(c *fiber.Ctx) error {
	var req SendRTMRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, "Invalid body", fiber.StatusBadRequest)
	}

	msg, err := SendRTMMessageService(req)
	if err != nil {
		return response.Error(c, "Failed to send RTM message", fiber.StatusInternalServerError)
	}

	return response.Success(c, msg, "Message sent", fiber.StatusOK)
}

func GetRTMHistory(c *fiber.Ctx) error {
	channelID := c.Query("channelId")
	limit := c.QueryInt("limit", 50)

	msgs, err := GetRTMMessages(channelID, limit)
	if err != nil {
		return response.Error(c, "Failed to load history", fiber.StatusInternalServerError)
	}

	return response.Success(c, msgs, "History loaded", fiber.StatusOK)
}

func GenerateRTMTokenHandler(c *fiber.Ctx) error {
	var req RTMTokenRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, "Invalid request", fiber.StatusBadRequest)
	}

	token, err := GenerateRTMToken(req.UserID)
	if err != nil {
		return response.Error(c, "Failed to generate token", fiber.StatusInternalServerError)
	}

	return response.Success(c, RTMTokenResponse{Token: token}, "RTM Token Generated", fiber.StatusOK)
}
