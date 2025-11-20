package response

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Code    int    `json:"code"`
	Errors  any    `json:"errors,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, data any, message string, code int) error {
	return c.Status(code).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
		Code:    code,
	})
}

func ValidationErrorResponse(c *fiber.Ctx, errors any, message string, code int) error {
	return c.Status(code).JSON(Response{
		Success: false,
		Message: message,
		Errors:  errors,
		Code:    code,
	})
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		msg = e.Message
	} else if err != nil {
		msg = err.Error()
	}

	log.Printf("[ERROR] %v", err)
	return c.Status(code).JSON(Response{
		Success: false,
		Message: msg,
		Code:    code,
	})
}
