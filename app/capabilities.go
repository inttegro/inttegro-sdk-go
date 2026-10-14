package app

import "time"

// CapabilityStatus describes the current availability of an app capability.
type CapabilityStatus string

const (
	CapabilityStatusEnabled       CapabilityStatus = "enabled"
	CapabilityStatusRestricted    CapabilityStatus = "restricted"
	CapabilityStatusPendingReview CapabilityStatus = "pending_review"
)

// CapabilityStatusReason explains why a capability is not enabled.
type CapabilityStatusReason string

const (
	CapabilityStatusReasonSetupIncomplete    CapabilityStatusReason = "setup_incomplete"
	CapabilityStatusReasonAccountUnderReview CapabilityStatusReason = "account_under_review"
	CapabilityStatusReasonAccountRestricted  CapabilityStatusReason = "account_restricted"
	CapabilityStatusReasonAccountInactive    CapabilityStatusReason = "account_inactive"
)

// CapabilityRequirement is a setup item that must be completed.
type CapabilityRequirement struct {
	Code        string     `json:"code"`
	Description string     `json:"description"`
	DueAt       *time.Time `json:"due_at"`
}

// CapabilityRequirements lists setup items that are currently due.
type CapabilityRequirements struct {
	CurrentlyDue []CapabilityRequirement `json:"currently_due"`
}

// Capability is the current state of one commerce capability.
type Capability struct {
	Status          CapabilityStatus        `json:"status"`
	StatusChangedAt time.Time               `json:"status_changed_at"`
	StatusReason    *CapabilityStatusReason `json:"status_reason,omitempty"`
	Requirements    *CapabilityRequirements `json:"requirements,omitempty"`
}

// Capabilities contains the payment and payout capabilities for an app.
type Capabilities struct {
	Payments Capability `json:"payments"`
	Payouts  Capability `json:"payouts"`
}
