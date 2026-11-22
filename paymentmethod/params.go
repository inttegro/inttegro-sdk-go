package paymentmethod

import (
	"github.com/inttegro/inttegro-sdk-go/v13/customdata"
	"github.com/inttegro/inttegro-sdk-go/v13/request"
)

type PageParams struct {
	CustomerID string `json:"customer_id,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
}

type ActionParams struct {
	PaymentMethodID string `json:"payment_method_id"`
}

// TokenizePaymentMethodParams saves a payment method for future use.
//
// Tokenization stores customer payment details securely for repeat charges.
// The customer owns the payment method—only they can delete it.
//
// Example:
//
//	params := paymentmethod.TokenizeParams{
//	    CustomerID: "cu_abc123",
//	    PaymentMethodData: paymentmethod.Data{
//	        Type: paymentmethod.TypeMobileMoney,
//	        MobileMoney: &paymentmethod.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    VerifyImmediately: inttegro.Bool(true),
//	}
type TokenizeParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// CustomerID is who owns this payment method (required).
	CustomerID string `json:"customer_id"`

	// PaymentMethodData contains the payment details to save (required).
	PaymentMethodData Data `json:"payment_method_data"`

	// VerifyImmediately triggers verification right after tokenization (optional).
	// If true, sends OTP immediately for customer to confirm ownership.
	// If false (default), must call Verify() separately before first use.
	VerifyImmediately *bool `json:"verify_immediately,omitempty"`
}

// VerifyPaymentMethodParams starts verification for a payment method.
//
// Sends an OTP to the payment method (phone number for mobile money)
// to confirm the customer owns it. Required before first use.
type VerifyParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// PaymentMethodID is the payment method to verify (required).
	PaymentMethodID string `json:"payment_method_id"`
}

// ConfirmPaymentMethodVerificationParams submits the OTP to complete verification.
type ConfirmVerificationParams struct {
	// PaymentMethodID is the payment method being verified (required).
	PaymentMethodID string `json:"payment_method_id"`

	// Token is the OTP from the customer (required).
	// Typically 4-6 digits sent via SMS.
	Token string `json:"token"`
}

// LookupPaymentMethodParams specifies which payment method to retrieve.
type LookupParams struct {
	// PaymentMethodID is the payment method identifier (required).
	// Starts with "pm_". Example: "pm_abc123def456"
	PaymentMethodID string `json:"payment_method_id"`
}

// UpdateParams changes the mutable fields of a saved payment method.
type UpdateParams struct {
	PaymentMethodID string            `json:"payment_method_id"`
	CustomData      *customdata.Patch `json:"custom_data,omitempty"`
	Active          *bool             `json:"active,omitempty"`
	Archived        *bool             `json:"archived,omitempty"`
	Owner           *UpdateOwner      `json:"owner,omitempty"`
}

// UpdateOwner contains owner fields that can be changed independently.
type UpdateOwner struct {
	Name    string         `json:"name,omitempty"`
	Address *UpdateAddress `json:"address,omitempty"`
}

// UpdateAddress contains mutable payment-method owner address fields.
type UpdateAddress struct {
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	Line1       string `json:"line1,omitempty"`
	Line2       string `json:"line2,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	PostCode    string `json:"post_code,omitempty"`
	Region      string `json:"region,omitempty"`
}

// DeletePaymentMethodParams permanently removes a payment method.
//
// Deleted payment methods cannot be restored. Customers can delete
// their own payment methods; merchants cannot force re-enablement.
type DeleteParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *request.Meta `json:"request_meta,omitempty"`

	// PaymentMethodID is the payment method to delete (required).
	PaymentMethodID string `json:"payment_method_id"`
}
