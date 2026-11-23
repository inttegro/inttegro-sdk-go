package financialaccount

import "time"

// PullPushConfig configures whether an account can send or receive funds.
//
// Pull configuration controls whether Inttegro can debit the account.
// Push configuration controls whether Inttegro can credit the account.
// Most payout destinations only need push enabled.
type PullPushConfig struct {
	// Enabled indicates whether the operation is allowed.
	// For payout destinations, PushConfiguration.Enabled should be true.
	// For payment sources, PullConfiguration.Enabled should be true.
	Enabled *bool `json:"enabled,omitempty"`

	// EnabledAt indicates when this configuration was enabled (read-only).
	EnabledAt *time.Time `json:"enabled_at,omitempty"`

	// Mandate contains mandate details for pull authorization (optional).
	Mandate *Mandate `json:"mandate,omitempty"`
}

// Mandate records the authorization that enabled account debits.
type Mandate struct {
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	IPAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
}
