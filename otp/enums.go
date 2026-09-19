package otp

type AlphabetType string

const (
	AlphabetTypeNumeric      AlphabetType = "numeric"
	AlphabetTypeAlpha        AlphabetType = "alpha"
	AlphabetTypeAlphanumeric AlphabetType = "alphanumeric"
)

type Purpose string

const (
	PurposeAccountCreation              Purpose = "account_creation"
	PurposeAccountRecovery              Purpose = "account_recovery"
	PurposeEmailVerification            Purpose = "email_verification"
	PurposeFinancialAccountVerification Purpose = "financial_account_verification"
	PurposePasswordReset                Purpose = "password_reset"
	PurposePaymentConfirmation          Purpose = "payment_confirmation"
	PurposePaymentMethodVerification    Purpose = "payment_method_verification"
	PurposePayoutConfirmation           Purpose = "payout_confirmation"
	PurposePhoneVerification            Purpose = "phone_verification"
	PurposeSensitiveAction              Purpose = "sensitive_action"
	PurposeSignIn                       Purpose = "sign_in"
	PurposeTransactionConfirmation      Purpose = "transaction_confirmation"
	PurposeUnspecified                  Purpose = "unspecified"
)

type Status string

const (
	StatusCanceled            Status = "canceled"
	StatusExpired             Status = "expired"
	StatusPending             Status = "pending"
	StatusPendingDelivery     Status = "pending_delivery"
	StatusPendingVerification Status = "pending_verification"
	StatusVerified            Status = "verified"
)

type TransmissionStatus string

const (
	TransmissionStatusDelivered TransmissionStatus = "delivered"
	TransmissionStatusFailed    TransmissionStatus = "failed"
	TransmissionStatusSubmitted TransmissionStatus = "submitted"
)

type VerificationVerdict string

const (
	VerificationVerdictFail VerificationVerdict = "fail"
	VerificationVerdictPass VerificationVerdict = "pass"
)
