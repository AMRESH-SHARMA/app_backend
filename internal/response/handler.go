package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

func Success(c *fiber.Ctx, data any, message string, code int) error {
	return c.Status(code).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ValidationError(c *fiber.Ctx, errors any, message string, code int) error {
	return c.Status(code).JSON(Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

func Error(c *fiber.Ctx, message string, code int) error {
	// code := fiber.StatusInternalServerError
	// msg := "Internal Server Error"

	// if e, ok := err.(*fiber.Error); ok {
	// 	code = e.Code
	// 	msg = e.Message
	// } else if err != nil {
	// 	msg = err.Error()
	// }
	// log.Printf("[ERROR] %v", err)
	return c.Status(code).JSON(Response{
		Success: false,
		Message: message,
	})
}
