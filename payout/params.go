package payout

import "time"

type ScheduleParams struct {
	DestinationID string     `json:"destination_id"`
	ExecuteAfter  *time.Time `json:"execute_after,omitempty"`
	MaxAmount     int64      `json:"max_amount,omitempty"`
	Reference     string     `json:"reference"`
}

// PayoutPageParams specifies pagination for listing payouts.
type PageParams struct {
	// PageNumber is the page to retrieve (optional, default: 1).
	// Pages are 1-indexed.
	PageNumber int `json:"page_number"`

	// PageSize is the number of payouts per page (optional, default: 256).
	// Maximum 256.
	PageSize int `json:"page_size,omitempty"`
}
