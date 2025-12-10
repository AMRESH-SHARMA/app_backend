package payment

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

/* ---------------- ENUMS ---------------- */

type OrderStatus string
type PaymentStatus string
type RefundStatus string

const (
	OrderPending   OrderStatus = "PENDING"
	OrderPaid      OrderStatus = "PAID"
	OrderCancelled OrderStatus = "CANCELLED"
	OrderRefunded  OrderStatus = "REFUNDED"

	PaymentPending        PaymentStatus = "PENDING"
	PaymentRequiresAction PaymentStatus = "REQUIRES_ACTION"
	PaymentSuccess        PaymentStatus = "SUCCESS"
	PaymentFailed         PaymentStatus = "FAILED"
	PaymentExpired        PaymentStatus = "EXPIRED"
	PaymentRefundPending  PaymentStatus = "REFUND_PENDING"
	PaymentRefunded       PaymentStatus = "REFUNDED"
	PaymentPartialRefund  PaymentStatus = "PARTIAL_REFUND"

	RefundPending RefundStatus = "PENDING"
	RefundSuccess RefundStatus = "SUCCESS"
	RefundFailed  RefundStatus = "FAILED"
)

/* ---------------- WALLET ---------------- */

type Wallet struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserId    uuid.UUID `gorm:"type:uuid;uniqueIndex"`
	Balance   int64     `gorm:"not null;default:0"`
	UpdatedAt time.Time
}

type WalletHistory struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserId    uuid.UUID `gorm:"type:uuid;index"`
	Amount    int64
	Type      string // CREDIT / DEBIT
	Reason    string
	CreatedAt time.Time
}

/* ---------------- ORDER ---------------- */

type Order struct {
	ID          uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserId      uuid.UUID   `gorm:"type:uuid;index"`
	Amount      int64
	Currency    string
	Status      OrderStatus
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
}

/* ---------------- PAYMENT ---------------- */

type Payment struct {
	ID               uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrderId          uuid.UUID     `gorm:"type:uuid;index"`
	UserId           uuid.UUID     `gorm:"type:uuid;index"`
	Gateway          string        `gorm:"size:50"`
	GatewayOrderId   string        `gorm:"size:100;index"`
	GatewayPaymentId string        `gorm:"size:100;index"`
	Amount           int64
	Currency         string
	Status           PaymentStatus
	Meta             datatypes.JSON
	FailureCode      *string
	FailureMessage   *string
	IdempotencyKey   string `gorm:"size:100;unique"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

/* ---------------- REFUND ---------------- */

type Refund struct {
	ID              uuid.UUID    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PaymentId       uuid.UUID    `gorm:"type:uuid;index"`
	GatewayRefundId string
	Amount          int64
	Status          RefundStatus
	Reason          string

	CreatedAt time.Time
	UpdatedAt time.Time
}

/* ---------------- WEBHOOK ---------------- */

type WebhookEvent struct {
	ID         uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Gateway    string        `gorm:"size:50"`
	EventId    string        `gorm:"size:200;uniqueIndex"`
	RawPayload datatypes.JSON
	Processed  bool
}
