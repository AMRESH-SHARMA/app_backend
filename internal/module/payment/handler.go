package payment

import (
	"app_backend/internal/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PaymentController struct {
	s Service
}

func NewPaymentController(s Service) *PaymentController {
	return &PaymentController{s}
}

/* WALLET */

func (c *PaymentController) GetRechargeOptions(ctx *fiber.Ctx) error {
	return response.Success(ctx, c.s.GetRechargeOptions(), "Fetched", 200)
}

func (c *PaymentController) GetBalance(ctx *fiber.Ctx) error {
	userId, _ := uuid.Parse(ctx.Params("id"))
	b, _ := c.s.GetBalance(ctx.Context(), userId)
	return response.Success(ctx, fiber.Map{"balance": b}, "Fetched", 200)
}

func (c *PaymentController) GetHistory(ctx *fiber.Ctx) error {
	userId, _ := uuid.Parse(ctx.Params("id"))
	h, _ := c.s.GetHistory(ctx.Context(), userId)
	return response.Success(ctx, h, "Fetched", 200)
}

/* ORDERS */

func (c *PaymentController) CreateOrder(ctx *fiber.Ctx) error {
	var req CreateOrderRequest
	ctx.BodyParser(&req)
	userId := uuid.New() // replace with JWT user
	res, err := c.s.CreateOrder(ctx.Context(), userId, req)
	if err != nil { return response.Error(ctx, err.Error(), 400) }
	return response.Success(ctx, res, "Order created", 200)
}

/* PAYMENT INTENT */

func (c *PaymentController) CreatePaymentIntent(ctx *fiber.Ctx) error {
	var req CreatePaymentIntentRequest
	ctx.BodyParser(&req)
	userId := uuid.New() // replace with JWT
	res, err := c.s.CreatePaymentIntent(ctx.Context(), userId, req)
	if err != nil { return response.Error(ctx, err.Error(), 400) }
	return response.Success(ctx, res, "Intent created", 200)
}

/* WEBHOOK */

func (c *PaymentController) Webhook(ctx *fiber.Ctx) error {
	var payload WebhookPayload
	ctx.BodyParser(&payload)
	c.s.Webhook(ctx.Context(), payload)
	return ctx.JSON(fiber.Map{"received": true})
}

/* ---------------- REFUND ---------------- */

func (c *PaymentController) CreateRefund(ctx *fiber.Ctx) error {
	var req CreateRefundRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.ValidationError(ctx, err.Error(), "Invalid request", fiber.StatusBadRequest)
	}

	res, err := c.s.CreateRefund(ctx.Context(), req)
	if err != nil {
		return response.Error(ctx, err.Error(), fiber.StatusBadRequest)
	}

	return response.Success(ctx, res, "Refund processed", fiber.StatusOK)
}