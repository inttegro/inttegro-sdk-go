package payout

// Destinations assigns each supported payout currency to a financial account.
type Destinations struct {
	// GHS is the financial account that receives Ghana cedi payouts.
	// Set it to an empty string when removing the assignment.
	GHS string `json:"ghs,omitempty"`
}
