package payout

// Page is one page of payouts returned by the API.
type Page struct {
	// Number is the returned 1-based page number.
	Number int `json:"number"`

	// Size is the number of payouts returned in this page.
	Size int `json:"size"`

	// Payouts contains the payouts in newest-first order when present.
	Payouts []Payout `json:"payouts,omitempty"`
}
