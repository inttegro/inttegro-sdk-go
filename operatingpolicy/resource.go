package operatingpolicy

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

// BTAgingSpec is the platform-resolved payout aging configuration.
type BTAgingSpec struct {
	ID    string `json:"id"`
	TPlus string `json:"t_plus"`
}

// BTAgingSpecReference selects a platform-owned payout aging configuration.
type BTAgingSpecReference struct {
	ID string `json:"id"`
}

// PaymentMethodCapabilities is a concrete payment-method configuration.
type PaymentMethodCapabilities struct {
	Enabled     bool `json:"enabled"`
	ConfirmsUse bool `json:"confirms_use"`
}

// PaymentMethodsCapabilities configures supported payment methods.
type PaymentMethodsCapabilities struct {
	MobileMoney *PaymentMethodCapabilities `json:"mobile_money,omitempty"`
	BankAccount *PaymentMethodCapabilities `json:"bank_account,omitempty"`
	Motito      *PaymentMethodCapabilities `json:"motito,omitempty"`
	Card        *PaymentMethodCapabilities `json:"card,omitempty"`
}

// PaymentsCapabilities is the resolved payment configuration.
type PaymentsCapabilities struct {
	Currencies []string `json:"currencies,omitempty"`
	Limits     *Limits  `json:"limits,omitempty"`
}

// PayoutsCapabilities is the resolved payout configuration.
type PayoutsCapabilities struct {
	Enabled     *bool        `json:"enabled,omitempty"`
	Limits      *Limits      `json:"limits,omitempty"`
	BTAgingSpec *BTAgingSpec `json:"bt_aging_spec,omitempty"`
}

// Capabilities is a resolved concrete operating configuration.
type Capabilities struct {
	Payments       *PaymentsCapabilities       `json:"payments,omitempty"`
	Payouts        *PayoutsCapabilities        `json:"payouts,omitempty"`
	PaymentMethods *PaymentMethodsCapabilities `json:"payment_methods,omitempty"`
}

// PaymentsSelection selects payment settings for a policy input.
type PaymentsSelection struct {
	Currencies []string `json:"currencies,omitempty"`
	Limits     *Limits  `json:"limits,omitempty"`
}

// PayoutsSelection selects payout settings for a policy input.
type PayoutsSelection struct {
	Enabled     *bool                 `json:"enabled,omitempty"`
	Limits      *Limits               `json:"limits,omitempty"`
	BTAgingSpec *BTAgingSpecReference `json:"bt_aging_spec,omitempty"`
}

// CapabilitiesSelection is the input form used to select concrete capabilities.
type CapabilitiesSelection struct {
	Payments       *PaymentsSelection          `json:"payments,omitempty"`
	Payouts        *PayoutsSelection           `json:"payouts,omitempty"`
	PaymentMethods *PaymentMethodsCapabilities `json:"payment_methods,omitempty"`
}

// Selection chooses either a named template or inline capabilities.
// Set exactly one of ID or Capabilities.
type Selection struct {
	ID           string                 `json:"id,omitempty"`
	Capabilities *CapabilitiesSelection `json:"capabilities,omitempty"`
}

// Source identifies how an application operating-policy snapshot was selected.
type Source struct {
	Type    string `json:"type"`
	ID      string `json:"id,omitempty"`
	Version int    `json:"version,omitempty"`
}

// Snapshot is the immutable operating policy applied to one application.
type Snapshot struct {
	Capabilities Capabilities `json:"capabilities"`
	Source       Source       `json:"source"`
}

// OperatingPolicy is one immutable version of an app-owned template.
type OperatingPolicy struct {
	ID                    string       `json:"id"`
	Version               int          `json:"version"`
	ActiveVersion         int          `json:"active_version"`
	Name                  string       `json:"name"`
	Description           string       `json:"description,omitempty"`
	Capabilities          Capabilities `json:"capabilities"`
	ChildrenPolicyID      string       `json:"children_policy_id"`
	ChildrenPolicyVersion int          `json:"children_policy_version"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             time.Time    `json:"updated_at"`
	VersionCreatedAt      time.Time    `json:"version_created_at"`
}

// Page contains immutable versions of one operating policy ordered newest first.
type Page struct {
	Number   int               `json:"number"`
	Size     int               `json:"size"`
	Count    int               `json:"count"`
	Total    int               `json:"total"`
	HasMore  bool              `json:"has_more"`
	Versions []OperatingPolicy `json:"versions"`
}
