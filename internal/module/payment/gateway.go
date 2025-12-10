package payment

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

/* Interface used by Razorpay, PhonePe, Stripe, Mock */

type PaymentGateway interface {
	CreateOrder(ctx context.Context, p *Payment) (*GatewayOrderResponse, error)
	Refund(ctx context.Context, p *Payment, amount int64) (*GatewayRefundResponse, error)
}

/* Response models */

type GatewayOrderResponse struct {
	GatewayOrderId string      `json:"gateway_order_id"`
	Meta           interface{} `json:"meta,omitempty"`
}

type GatewayRefundResponse struct {
	GatewayRefundId string `json:"gateway_refund_id"`
	Status          string `json:"status"`
}

/* ---------------- MOCK Gateway ---------------- */

type MockGateway struct{}

func NewMockGateway() PaymentGateway { return &MockGateway{} }

func (m *MockGateway) CreateOrder(ctx context.Context, p *Payment) (*GatewayOrderResponse, error) {
	return &GatewayOrderResponse{
		GatewayOrderId: "mock_order_" + p.ID.String(),
		Meta:           map[string]string{"msg": "Mock gateway"},
	}, nil
}

func (m *MockGateway) Refund(ctx context.Context, p *Payment, amt int64) (*GatewayRefundResponse, error) {
	return &GatewayRefundResponse{
		GatewayRefundId: fmt.Sprintf("mock_r_%s", uuid.NewString()),
		Status:          "SUCCESS",
	}, nil
}
