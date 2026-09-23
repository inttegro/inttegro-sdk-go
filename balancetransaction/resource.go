package balancetransaction

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v10/money"
	"github.com/zebodotdev/inttegro-sdk-go/v10/payout"
)

// BalanceTransaction represents a merchant balance entry caused by a payment or
// refund. Exactly one of PaymentID and RefundID is present, matching Type.
type BalanceTransaction struct {
	// ID is the unique balance transaction identifier (read-only).
	// Starts with "bt_". Example: "bt_abc123def456"
	ID string `json:"id"`

	// Type identifies the semantic source, not the accounting direction.
	Type Type `json:"type"`

	// PaymentID is present only when Type is payment.
	PaymentID string `json:"payment_id,omitempty"`

	// RefundID is present only when Type is refund.
	RefundID string `json:"refund_id,omitempty"`

	// PayoutID identifies the legacy payout that claimed this whole transaction.
	// Deprecated: inspect Allocations because one transaction can fund many payouts.
	PayoutID string `json:"payout_id,omitempty"`

	// OrderID is the strongly referenced source order ID.
	OrderID string `json:"order_id"`

	// Amount is the transaction amount in the public money shape.
	Amount money.Amount `json:"amount"`

	// Allocations contains every current pending and completed allocation of a
	// payment transaction. Released allocations are omitted.
	Allocations []Allocation `json:"allocations,omitempty"`

	// AvailableAmount is the amount still available for a refund or payout.
	AvailableAmount *money.Amount `json:"available_amount,omitempty"`

	// PendingAmount is the amount reserved by unresolved refunds and payouts.
	PendingAmount *money.Amount `json:"pending_amount,omitempty"`

	// SpentAmount is the amount permanently consumed by completed allocations.
	SpentAmount *money.Amount `json:"spent_amount,omitempty"`

	// AvailableAt is when funds become eligible for payout (ISO 8601, read-only).
	AvailableAt *time.Time `json:"available_at,omitempty"`

	// ClaimedAt is the legacy whole-transaction payout claim time.
	// Deprecated: inspect Allocations for current payout participation.
	ClaimedAt *time.Time `json:"claimed_at,omitempty"`

	// PaidAt is the legacy whole-transaction payout completion time.
	// Deprecated: inspect completed Allocations for consumed amounts.
	PaidAt *time.Time `json:"paid_at,omitempty"`

	// CreatedAt is when the balance transaction was created (ISO 8601).
	CreatedAt time.Time `json:"created_at"`

	// PayoutConfiguration is the routing used when this transaction is paid out.
	PayoutConfiguration *payout.Configuration `json:"payout_configuration,omitempty"`
}

// Allocation is a caller-safe record of part of a payment balance transaction
// being assigned to a refund or payout. Exactly one of Refund and Payout is
// present, matching Type.
type Allocation struct {
	ID          string           `json:"id"`
	Type        AllocationType   `json:"type"`
	Status      AllocationStatus `json:"status"`
	Refund      *AllocationUse   `json:"refund,omitempty"`
	Payout      *AllocationUse   `json:"payout,omitempty"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	CompletedAt *time.Time       `json:"completed_at,omitempty"`
}

// AllocationUse identifies the refund or payout and the amount it consumed
// from this balance transaction.
type AllocationUse struct {
	ID     string       `json:"id"`
	Amount money.Amount `json:"amount"`
}

// SourceID returns the matching strong source reference. It returns false for
// incomplete or contradictory transaction values.
func (t BalanceTransaction) SourceID() (string, bool) {
	switch t.Type {
	case TypePayment:
		if t.PaymentID != "" && t.RefundID == "" {
			return t.PaymentID, true
		}
	case TypeRefund:
		if t.RefundID != "" && t.PaymentID == "" {
			return t.RefundID, true
		}
	}
	return "", false
}
