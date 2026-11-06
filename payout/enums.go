package payout

type Status string

const (
	StatusInitialized Status = "initialized"
	StatusScheduled   Status = "scheduled"
	StatusProcessing  Status = "processing"
	StatusExecuting   Status = "executing"
	StatusSucceeded   Status = "succeeded"
	StatusFailed      Status = "failed"
	StatusCanceled    Status = "canceled"
)

// BalanceTransactionStatus describes what happened to a balance transaction's
// contribution to a payout.
type BalanceTransactionStatus string

const (
	BalanceTransactionStatusReserved  BalanceTransactionStatus = "reserved"
	BalanceTransactionStatusCompleted BalanceTransactionStatus = "completed"
	BalanceTransactionStatusReleased  BalanceTransactionStatus = "released"
)
