package balancetransaction

type Type string

const (
	TypePayment Type = "payment"
	TypeRefund  Type = "refund"
)

// AllocationType identifies what consumed part of a payment balance transaction.
type AllocationType string

const (
	AllocationTypePayout AllocationType = "payout"
	AllocationTypeRefund AllocationType = "refund"
)

// AllocationStatus describes the caller-visible effect of an allocation.
// Internal reservation and reconciliation states are intentionally collapsed.
type AllocationStatus string

const (
	AllocationStatusPending   AllocationStatus = "pending"
	AllocationStatusCompleted AllocationStatus = "completed"
)
