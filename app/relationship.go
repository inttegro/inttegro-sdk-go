package app

import "time"

// AppRelationshipPolicy describes who manages a child app relationship.
type RelationshipPolicy struct {
	// ChildStanding is the initial standing assigned to the child app.
	ChildStanding string `json:"child_standing,omitempty"`

	// Management controls who can manage the child app and its resources.
	Management ManagementRole `json:"management,omitempty"`

	// Credentials controls who can create, rotate, or disable the child app's API keys.
	Credentials CredentialOwner `json:"credentials,omitempty"`
}

// AppRelationship is the placement relationship receipt for a child app.
type Relationship struct {
	ID        string             `json:"id"`
	Kind      RelationshipKind   `json:"kind"`
	Parent    RelationshipParent `json:"parent"`
	Status    RelationshipStatus `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
}

// RelationshipParent identifies the app that directly contains a child app.
type RelationshipParent struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Alias string `json:"alias,omitempty"`
}
