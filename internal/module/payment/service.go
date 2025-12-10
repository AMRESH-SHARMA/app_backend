package payment

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Service interface {
	GetRechargeOptions() []RechargeOption
	GetBalance(ctx context.Context, userId uuid.UUID) (int64, error)
	GetHistory(ctx context.Context, userId uuid.UUID) ([]WalletHistory, error)

	CreateOrder(ctx context.Context, userId uuid.UUID, req CreateOrderRequest) (*CreateOrderResponse, error)
	CreatePaymentIntent(ctx context.Context, userId uuid.UUID, req CreatePaymentIntentRequest) (*CreatePaymentIntentResponse, error)
	Webhook(ctx context.Context, payload WebhookPayload) error
	CreateRefund(ctx context.Context, req CreateRefundRequest) (*CreateRefundResponse, error)
}

type service struct {
	repo     Repository
	gateways map[string]PaymentGateway
}

func NewService(r Repository, gw map[string]PaymentGateway) Service {
	return &service{repo: r, gateways: gw}
}

/* ------------- WALLET ------------- */

func (s *service) GetRechargeOptions() []RechargeOption {
	return []RechargeOption{
		{Amount: 100, Bonus: 0, Currency: "INR"},
		{Amount: 200, Bonus: 10, Currency: "INR"},
		{Amount: 500, Bonus: 40, Currency: "INR"},
		{Amount: 1000, Bonus: 100, Currency: "INR"},
	}
}

func (s *service) GetBalance(ctx context.Context, userId uuid.UUID) (int64, error) {
	w, err := s.repo.GetOrCreateWallet(ctx, userId)
	if err != nil { return 0, err }
	return w.Balance, nil
}

func (s *service) GetHistory(ctx context.Context, userId uuid.UUID) ([]WalletHistory, error) {
	w, err := s.repo.GetOrCreateWallet(ctx, userId)
	if err != nil { return nil, err }

	var history []WalletHistory
	_ = s.repo.(*repo).db.WithContext(ctx).Where("user_id = ?", w.UserId).
		Order("created_at desc").
		Find(&history)

	return history, nil
}

/* ------------- ORDER ------------- */

func (s *service) CreateOrder(ctx context.Context, userId uuid.UUID, req CreateOrderRequest) (*CreateOrderResponse, error) {
	if req.Amount <= 0 {
		return nil, errors.New("amount must be > 0")
	}

	order := &Order{
		ID:          uuid.New(),
		UserId:      userId,
		Amount:      req.Amount,
		Currency:    req.Currency,
		Status:      OrderPending,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	return &CreateOrderResponse{
		OrderId: order.ID,
		Status:  string(order.Status),
	}, nil
}

/* ------------- PAYMENT INTENT ------------- */

func (s *service) CreatePaymentIntent(ctx context.Context, userId uuid.UUID, req CreatePaymentIntentRequest) (*CreatePaymentIntentResponse, error) {

	order, err := s.repo.GetOrder(ctx, req.OrderId)
	if err != nil {
		return nil, errors.New("order not found")
	}

	gw, exists := s.gateways[req.Gateway]
	if !exists {
		return nil, errors.New("gateway not supported")
	}

	p := &Payment{
		ID:             uuid.New(),
		OrderId:        order.ID,
		UserId:         order.UserId,
		Gateway:        req.Gateway,
		Amount:         order.Amount,
		Currency:       order.Currency,
		Status:         PaymentPending,
		IdempotencyKey: uuid.NewString(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	s.repo.CreatePayment(ctx, p)

	// Call gateway
	gwResp, err := gw.CreateOrder(ctx, p)
	if err != nil {
		return nil, err
	}

	p.GatewayOrderId = gwResp.GatewayOrderId
	p.Status = PaymentRequiresAction
	s.repo.UpdatePayment(ctx, p)

	return &CreatePaymentIntentResponse{
		PaymentId:      p.ID,
		OrderId:        p.OrderId,
		Gateway:        p.Gateway,
		GatewayOrderId: p.GatewayOrderId,
		Amount:         p.Amount,
		Status:         string(p.Status),
	}, nil
}

/* ------------- WEBHOOK (Source of truth) ------------- */

func (s *service) Webhook(ctx context.Context, payload WebhookPayload) error {

	raw, _ := json.Marshal(payload)
	s.repo.CreateWebhookEvent(ctx, &WebhookEvent{
		ID:         uuid.New(),
		Gateway:    payload.Gateway,
		EventId:    payload.EventId,
		RawPayload: datatypes.JSON(raw),
		Processed:  true,
	})

	p, err := s.repo.GetPaymentByGatewayPaymentId(ctx, payload.Gateway, payload.GatewayPaymentId)
	if err != nil { return nil } // ignore unknown events

	switch payload.EventType {

	case "PAYMENT_SUCCESS":
		p.Status = PaymentSuccess
		p.GatewayPaymentId = payload.GatewayPaymentId
		s.repo.UpdatePayment(ctx, p)

		// Update order
		o, _ := s.repo.GetOrder(ctx, p.OrderId)
		o.Status = OrderPaid
		s.repo.UpdateOrder(ctx, o)

		// Credit wallet
		w, _ := s.repo.GetOrCreateWallet(ctx, p.UserId)
		w.Balance += payload.Amount
		w.UpdatedAt = time.Now()
		s.repo.UpdateWallet(ctx, w)

		s.repo.AddWalletHistory(ctx, &WalletHistory{
			ID:        uuid.New(),
			UserId:    p.UserId,
			Amount:    payload.Amount,
			Type:      "CREDIT",
			Reason:    "Wallet recharge",
			CreatedAt: time.Now(),
		})

	case "PAYMENT_FAILED":
		p.Status = PaymentFailed
		s.repo.UpdatePayment(ctx, p)

	case "PAYMENT_EXPIRED":
		p.Status = PaymentExpired
		s.repo.UpdatePayment(ctx, p)
	}

	return nil
}

/* ------------- REFUND ------------- */

func (s *service) CreateRefund(ctx context.Context, req CreateRefundRequest) (*CreateRefundResponse, error) {

	p, err := s.repo.(*repo).GetPaymentById(ctx, req.PaymentId)
	if err != nil { return nil, errors.New("payment not found") }

	gw := s.gateways[p.Gateway]

	gwResp, err := gw.Refund(ctx, p, req.Amount)
	if err != nil { return nil, err }

	ref := &Refund{
		ID:              uuid.New(),
		PaymentId:       p.ID,
		GatewayRefundId: gwResp.GatewayRefundId,
		Amount:          req.Amount,
		Status:          RefundSuccess,
		Reason:          req.Reason,
		CreatedAt:       time.Now(),
	}

	s.repo.CreateRefund(ctx, ref)

	return &CreateRefundResponse{
		RefundId: ref.ID,
		Amount:   ref.Amount,
		Status:   string(ref.Status),
	}, nil
}
