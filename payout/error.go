package payout

import "time"

// Error contains the public details of a payout execution failure.
type Error struct {
	// Cause explains the underlying failure in caller-safe terms.
	Cause string `json:"cause"`

	// Message is the human-readable failure message.
	Message string `json:"message"`

	// OccurredAt is when the failure occurred.
	OccurredAt time.Time `json:"occurred_at"`

	// Type is the stable machine-readable failure category.
	Type string `json:"type"`
}
