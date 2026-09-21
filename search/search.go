// Package search contains request and response types shared by resource search endpoints.
package search

type Operator string

const (
	OperatorEqual Operator = "eq"
	OperatorIn    Operator = "in"
)

type SortField string

const (
	SortRelevance   SortField = "relevance"
	SortUpdatedAt   SortField = "updated_at"
	SortPublishedAt SortField = "published_at"
)

type SortDirection string

const (
	SortAscending  SortDirection = "asc"
	SortDescending SortDirection = "desc"
)

type ResourceType string

const (
	ResourceCustomer         ResourceType = "customer"
	ResourceFinancialAccount ResourceType = "financial_account"
	ResourceOrder            ResourceType = "order"
	ResourcePayout           ResourceType = "payout"
	ResourceProduct          ResourceType = "product"
)

type TotalRelation string

const (
	TotalExact      TotalRelation = "exact"
	TotalLowerBound TotalRelation = "lower_bound"
)

type FreshnessState string

const (
	FreshnessCurrent     FreshnessState = "current"
	FreshnessDelayed     FreshnessState = "delayed"
	FreshnessPartial     FreshnessState = "partial"
	FreshnessUnknown     FreshnessState = "unknown"
	FreshnessUnavailable FreshnessState = "unavailable"
)

type Filter struct {
	Field    string   `json:"field"`
	Operator Operator `json:"operator"`
	Values   []string `json:"values"`
}

type Facet struct {
	Field string `json:"field"`
	Limit int    `json:"limit,omitempty"`
}

type Sort struct {
	Field     SortField     `json:"field"`
	Direction SortDirection `json:"direction"`
}

// Params describes one resource-local search.
type Params struct {
	Text     string   `json:"text,omitempty"`
	Filters  []Filter `json:"filters,omitempty"`
	Facets   []Facet  `json:"facets,omitempty"`
	Sort     *Sort    `json:"sort,omitempty"`
	PageSize int      `json:"page_size,omitempty"`
	Cursor   string   `json:"cursor,omitempty"`
}
