package payment

import "github.com/google/uuid"

type RechargeOption struct {
	Amount   int64  `json:"amount"`
	Bonus    int64  `json:"bonus"`
	Currency string `json:"currency"`
}

type CreateOrderRequest struct {
	Amount      int64  `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description"`
}

type CreateOrderResponse struct {
	OrderId uuid.UUID `json:"order_id"`
	Status  string    `json:"status"`
}

type CreatePaymentIntentRequest struct {
	OrderId uuid.UUID `json:"order_id"`
	Gateway string    `json:"gateway"` // mock | razorpay | phonepe | stripe
}

type CreatePaymentIntentResponse struct {
	PaymentId      uuid.UUID `json:"payment_id"`
	OrderId        uuid.UUID `json:"order_id"`
	Gateway        string    `json:"gateway"`
	GatewayOrderId string    `json:"gateway_order_id"`
	Amount         int64     `json:"amount"`
	Status         string    `json:"status"`
}

type WebhookPayload struct {
	EventId          string `json:"event_id"`
	EventType        string `json:"event_type"`
	Gateway          string `json:"gateway"`
	GatewayPaymentId string `json:"gateway_payment_id"`
	GatewayOrderId   string `json:"gateway_order_id"`
	Amount           int64  `json:"amount"`
	Currency         string `json:"currency"`
}

type CreateRefundRequest struct {
	PaymentId uuid.UUID `json:"payment_id"`
	Amount    int64     `json:"amount"`
	Reason    string    `json:"reason"`
}

type CreateRefundResponse struct {
	RefundId uuid.UUID `json:"refund_id"`
	Amount   int64     `json:"amount"`
	Status   string    `json:"status"`
}
