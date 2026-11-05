package childrenpolicy

import "time"

// AmountRange is an inclusive minor-unit amount range.
type AmountRange struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

// Limits contains inclusive configuration ranges.
type Limits struct {
	Amount AmountRange `json:"amount"`
}

// BTAgingSpec is one permitted payout balance-transaction aging option.
type BTAgingSpec struct {
	PublicID string `json:"public_id"`
	TPlus    string `json:"t_plus"`
}

// PaymentMethodCapabilities lists permitted values for one payment method.
type PaymentMethodCapabilities struct {
	Enabled     []bool `json:"enabled"`
	ConfirmsUse []bool `json:"confirms_use"`
}

// PaymentMethodsCapabilities constrains the supported payment methods.
type PaymentMethodsCapabilities struct {
	MobileMoney *PaymentMethodCapabilities `json:"mobile_money,omitempty"`
	BankAccount *PaymentMethodCapabilities `json:"bank_account,omitempty"`
	Motito      *PaymentMethodCapabilities `json:"motito,omitempty"`
	Card        *PaymentMethodCapabilities `json:"card,omitempty"`
}

// PaymentsCapabilities constrains payment settings a child may select.
type PaymentsCapabilities struct {
	Currencies []string `json:"currencies,omitempty"`
	Limits     *Limits  `json:"limits,omitempty"`
}

// PayoutsCapabilities constrains payout settings a child may select.
type PayoutsCapabilities struct {
	Enabled     []bool        `json:"enabled,omitempty"`
	Limits      *Limits       `json:"limits,omitempty"`
	BTAgingSpec []BTAgingSpec `json:"bt_aging_spec,omitempty"`
}

// Capabilities is the authority granted by a children policy.
type Capabilities struct {
	Payments       *PaymentsCapabilities       `json:"payments,omitempty"`
	Payouts        *PayoutsCapabilities        `json:"payouts,omitempty"`
	PaymentMethods *PaymentMethodsCapabilities `json:"payment_methods,omitempty"`
}

// ChildrenPolicy is the active immutable version assigned to an application.
type ChildrenPolicy struct {
	ID               string       `json:"id"`
	Version          int          `json:"version"`
	ActiveVersion    int          `json:"active_version"`
	Name             string       `json:"name"`
	Description      string       `json:"description,omitempty"`
	Capabilities     Capabilities `json:"capabilities"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	VersionCreatedAt time.Time    `json:"version_created_at"`
}
