package payout

// FailureReason is a stable, caller-safe reason that a payout failed.
type FailureReason string

const (
	FailureReasonProviderDeclined       FailureReason = "provider_declined"
	FailureReasonDeliveryFailed         FailureReason = "delivery_failed"
	FailureReasonTemporarilyUnavailable FailureReason = "temporarily_unavailable"
	FailureReasonUnknown                FailureReason = "unknown"
)

// Failure contains sanitized terminal payout failure information.
type Failure struct {
	Detail    string        `json:"detail"`
	Reason    FailureReason `json:"reason"`
	Retryable bool          `json:"retryable"`
}
