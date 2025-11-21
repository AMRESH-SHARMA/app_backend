package payment

import (
	"app_backend/internal/response"

	"github.com/gofiber/fiber/v2"
)

func GetRechargeOptions(c *fiber.Ctx) error {
	return response.Success(c, fiber.Map{}, "Data", fiber.StatusOK)
}

func GetBalance(c *fiber.Ctx) error {
	return response.Success(c, fiber.Map{}, "Data", fiber.StatusOK)
}

func AddBalance(c *fiber.Ctx) error {
	// id := c.Params("id")
	return response.Success(c, fiber.Map{}, "Data", fiber.StatusOK)
}

func GetHistory(c *fiber.Ctx) error {
	return response.Success(c, fiber.Map{}, "Data", fiber.StatusOK)
}
