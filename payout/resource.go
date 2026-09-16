package payout

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v9/money"
)

// Payout represents a settlement transfer to your bank or mobile money account.
//
// Payouts move funds from your Inttegro balance to your financial accounts.
// Each payout contains one or more balance transactions that have aged
// past the dispute window.
type Payout struct {
	// Amount is the amount transferred once payout execution begins.
	// It is absent while only the maximum payout amount is known.
	Amount *money.Amount `json:"amount,omitempty"`

	// BalanceTransactions contains the balance transaction IDs included in the payout.
	BalanceTransactions []string `json:"balance_transactions,omitempty"`

	// CanceledAt is when a scheduled payout was canceled.
	CanceledAt *time.Time `json:"canceled_at,omitempty"`

	// CustomData contains merchant-defined string metadata attached to the payout.
	CustomData *CustomData `json:"custom_data,omitempty"`

	// DestinationID identifies the financial account receiving the payout.
	DestinationID string `json:"destination_id"`

	// Error describes a public payout execution failure.
	Error *Error `json:"error,omitempty"`

	// ExecuteAfter is the earliest time at which payout execution may begin.
	ExecuteAfter time.Time `json:"execute_after"`

	// ExecutedBy identifies the actor that executed the payout.
	ExecutedBy string `json:"executed_by,omitempty"`

	// ExpectedAt is the estimated payout completion time.
	ExpectedAt *time.Time `json:"expected_at,omitempty"`

	// FailedAt is when the payout entered its unsuccessful terminal state.
	FailedAt *time.Time `json:"failed_at,omitempty"`

	// ID is the unique payout identifier.
	ID string `json:"id"`

	// InitiatedAt is when the payout was created.
	InitiatedAt time.Time `json:"initiated_at"`

	// InitiatedBy identifies the actor that initiated the payout.
	InitiatedBy string `json:"initiated_by,omitempty"`

	// MaxAmount is the maximum amount authorized for this payout.
	MaxAmount money.Amount `json:"max_amount"`

	// Reference is the merchant-supplied payout reference.
	Reference string `json:"reference,omitempty"`

	// ScheduleID identifies the schedule associated with the payout.
	ScheduleID string `json:"schedule_id,omitempty"`

	// ScheduledAt is when the payout was scheduled.
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`

	// ScheduledBy identifies the actor that scheduled the payout.
	ScheduledBy string `json:"scheduled_by,omitempty"`

	// SentAt is when the payout transfer was sent.
	SentAt *time.Time `json:"sent_at,omitempty"`

	// SourceID identifies the payout's source when one was recorded.
	SourceID string `json:"source_id,omitempty"`

	// Status is the payout's current lifecycle state.
	Status Status `json:"status"`

	// SucceededAt is when the payout completed successfully.
	SucceededAt *time.Time `json:"succeeded_at,omitempty"`
}
