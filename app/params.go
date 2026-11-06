package app

import "github.com/inttegro/inttegro-sdk-go/v13/operatingpolicy"

// CreateAppParams creates an Inttegro child app.
type CreateParams struct {
	// Name is the app's display name. It must not be empty after trimming.
	Name string `json:"name"`

	// Alias is an optional short label for dashboards, logs, or tooling.
	Alias string `json:"alias,omitempty"`

	// Description describes what the app is used for.
	Description string `json:"description,omitempty"`

	// LegalEntityType is a caller-defined legal-entity label.
	LegalEntityType string `json:"legal_entity_type,omitempty"`

	// PlacementParentApplicationID places the child under another app when separately authorized.
	PlacementParentApplicationID string `json:"placement_parent_application_id,omitempty"`

	// OperatingPolicy selects a parent-owned template or inline capabilities.
	OperatingPolicy *operatingpolicy.Selection `json:"operating_policy,omitempty"`

	// RelationshipPolicy customizes the direct parent-child relationship.
	RelationshipPolicy *RelationshipPolicy `json:"relationship_policy,omitempty"`
}

// UpdateAppParams changes mutable metadata on the authenticated app.
type UpdateParams struct {
	// Name changes the app's display name. It must not be empty after trimming when supplied.
	Name *string `json:"name,omitempty"`

	// Alias changes or clears the app's short label.
	Alias *string `json:"alias,omitempty"`

	// Description changes or clears the app description.
	Description *string `json:"description,omitempty"`

	// LegalEntityType changes or clears the caller-defined legal-entity label.
	LegalEntityType *string `json:"legal_entity_type,omitempty"`

	// OperatingPolicy replaces the app's immutable operating-policy snapshot.
	OperatingPolicy *operatingpolicy.Selection `json:"operating_policy,omitempty"`
}
