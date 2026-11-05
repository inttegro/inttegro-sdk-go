package operatingpolicy

// CreateParams creates the first version of an operating policy.
type CreateParams struct {
	Name         string                `json:"name"`
	Description  string                `json:"description,omitempty"`
	Capabilities CapabilitiesSelection `json:"capabilities"`
}

// LookupParams identifies an operating policy owned by the authenticated app.
type LookupParams struct {
	OperatingPolicyID string `json:"operating_policy_id"`
}

// UpdateParams creates and activates a new immutable policy version.
type UpdateParams struct {
	OperatingPolicyID string                 `json:"operating_policy_id"`
	ExpectedVersion   int                    `json:"expected_version"`
	Name              *string                `json:"name,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Capabilities      *CapabilitiesSelection `json:"capabilities,omitempty"`
}
