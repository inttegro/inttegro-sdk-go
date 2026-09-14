package refund

// FailureReason is a stable, caller-safe reason that terminal refund
// processing failed.
type FailureReason string

const (
	FailureReasonInsufficientBalance              FailureReason = "insufficient_balance"
	FailureReasonOriginalPaymentMethodUnavailable FailureReason = "original_payment_method_unavailable"
	FailureReasonOriginalPaymentNotRefundable     FailureReason = "original_payment_not_refundable"
	FailureReasonRefundNotSupported               FailureReason = "refund_not_supported"
	FailureReasonAmountNotSupported               FailureReason = "amount_not_supported"
	FailureReasonRefundDeclined                   FailureReason = "refund_declined"
	FailureReasonRefundNotPermitted               FailureReason = "refund_not_permitted"
	FailureReasonTemporarilyUnavailable           FailureReason = "temporarily_unavailable"
	FailureReasonUnknown                          FailureReason = "unknown"
)

// Failure contains sanitized terminal failure information. Provider and
// internal diagnostics are never returned in this value.
type Failure struct {
	Reason    FailureReason `json:"reason"`
	Detail    string        `json:"detail"`
	Retryable bool          `json:"retryable"`
}
