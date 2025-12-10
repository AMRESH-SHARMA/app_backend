package payment

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	// Wallet
	GetOrCreateWallet(ctx context.Context, userId uuid.UUID) (*Wallet, error)
	UpdateWallet(ctx context.Context, wallet *Wallet) error
	AddWalletHistory(ctx context.Context, h *WalletHistory) error

	// Orders
	CreateOrder(ctx context.Context, o *Order) error
	GetOrder(ctx context.Context, id uuid.UUID) (*Order, error)
	UpdateOrder(ctx context.Context, o *Order) error

	// Payments
	CreatePayment(ctx context.Context, p *Payment) error
	UpdatePayment(ctx context.Context, p *Payment) error
	GetPaymentById(ctx context.Context, id uuid.UUID) (*Payment, error)
	GetPaymentByGatewayPaymentId(ctx context.Context, gateway, pid string) (*Payment, error)

	// Refund
	CreateRefund(ctx context.Context, r *Refund) error

	// Webhook
	CreateWebhookEvent(ctx context.Context, w *WebhookEvent) error
}

type repo struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

/* WALLETS */

func (r *repo) GetOrCreateWallet(ctx context.Context, userId uuid.UUID) (*Wallet, error) {
	var w Wallet
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).First(&w).Error
	if err == gorm.ErrRecordNotFound {
		w = Wallet{ID: uuid.New(), UserId: userId, Balance: 0}
		err = r.db.WithContext(ctx).Create(&w).Error
	}
	return &w, err
}

func (r *repo) UpdateWallet(ctx context.Context, w *Wallet) error {
	return r.db.WithContext(ctx).Save(w).Error
}

func (r *repo) AddWalletHistory(ctx context.Context, h *WalletHistory) error {
	return r.db.WithContext(ctx).Create(h).Error
}

/* ORDERS */

func (r *repo) CreateOrder(ctx context.Context, o *Order) error {
	return r.db.WithContext(ctx).Create(o).Error
}

func (r *repo) GetOrder(ctx context.Context, id uuid.UUID) (*Order, error) {
	var o Order
	err := r.db.WithContext(ctx).First(&o, "id = ?", id).Error
	return &o, err
}

func (r *repo) UpdateOrder(ctx context.Context, o *Order) error {
	return r.db.WithContext(ctx).Save(o).Error
}

/* PAYMENTS */

func (r *repo) CreatePayment(ctx context.Context, p *Payment) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *repo) UpdatePayment(ctx context.Context, p *Payment) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *repo) GetPaymentById(ctx context.Context, id uuid.UUID) (*Payment, error) {
	var p Payment
	err := r.db.WithContext(ctx).First(&p, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}


func (r *repo) GetPaymentByGatewayPaymentId(ctx context.Context, gateway, pid string) (*Payment, error) {
	var p Payment
	err := r.db.WithContext(ctx).First(&p,
		"gateway = ? AND gateway_payment_id = ?", gateway, pid).Error
	return &p, err
}

/* REFUNDS */

func (r *repo) CreateRefund(ctx context.Context, rf *Refund) error {
	return r.db.WithContext(ctx).Create(rf).Error
}

/* WEBHOOK */

func (r *repo) CreateWebhookEvent(ctx context.Context, w *WebhookEvent) error {
	return r.db.WithContext(ctx).Create(w).Error
}
