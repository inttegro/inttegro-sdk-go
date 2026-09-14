package payout

// SettingsLookup contains the complete payout configuration returned by
// Service.Settings.
type SettingsLookup struct {
	// Destinations contains the configured account for each supported currency.
	Destinations Destinations `json:"destinations"`

	// FXEnabled reports whether payout currency conversion is enabled.
	// It is absent when the setting has not been configured.
	FXEnabled *bool `json:"fx_enabled,omitempty"`

	// Schedule is the active payout schedule, when configured.
	Schedule *SettingsLookupSchedule `json:"schedule,omitempty"`
}

// SettingsMutation contains the fields returned after a payout settings
// mutation. Endpoints return only the fields relevant to that mutation.
type SettingsMutation struct {
	// Destinations contains the updated destination assignments, when returned.
	Destinations *Destinations `json:"destinations,omitempty"`

	// FXEnabled contains the updated currency-conversion state, when returned.
	FXEnabled *bool `json:"fx_enabled,omitempty"`

	// ID is the payout settings identifier, when returned.
	ID string `json:"id,omitempty"`

	// Schedule is the updated payout schedule, when returned.
	Schedule *SettingsMutationSchedule `json:"schedule,omitempty"`
}

// PayoutConfiguration describes payout routing and FX settings for a payment or balance transaction.
type Configuration struct {
	// EnableFX indicates whether FX conversion is enabled for this payout.
	EnableFX bool `json:"enable_fx"`

	// Destination specifies the financial account receiving the payout.
	Destination Destination `json:"destination"`
}

// PayoutDestination identifies the payout financial account.
type Destination struct {
	// FinancialAccountID is the ID of the destination financial account.
	FinancialAccountID string `json:"financial_account_id"`
}
