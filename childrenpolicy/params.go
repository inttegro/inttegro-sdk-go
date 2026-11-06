package childrenpolicy

// LookupParams optionally selects one exact children-policy version.
// Version zero is omitted and resolves to the latest version.
type LookupParams struct {
	Version int `json:"version,omitempty"`
}

// PageParams controls children-policy version pagination.
type PageParams struct {
	Number int `json:"number,omitempty"`
	Size   int `json:"size,omitempty"`
}
