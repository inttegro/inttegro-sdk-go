package search

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v10/money"
)

type Total struct {
	Value    int64         `json:"value"`
	Relation TotalRelation `json:"relation"`
}

type ResourceTotal struct {
	ResourceType ResourceType  `json:"resource_type"`
	Value        int64         `json:"value"`
	Relation     TotalRelation `json:"relation"`
}

type ResourceReference struct {
	Type ResourceType `json:"type"`
	ID   string       `json:"id"`
}

type Result struct {
	Resource     ResourceReference `json:"resource"`
	Title        string            `json:"title"`
	Summary      string            `json:"summary,omitempty"`
	Status       string            `json:"status,omitempty"`
	CustomerName string            `json:"customer_name,omitempty"`
	Amount       *money.Amount     `json:"amount,omitempty"`
	URL          string            `json:"url,omitempty"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

type FacetBucket struct {
	Value string `json:"value"`
	Count int64  `json:"count"`
}

type FacetResult struct {
	Field   string        `json:"field"`
	Buckets []FacetBucket `json:"buckets"`
}

type ResourceFreshness struct {
	ResourceType  ResourceType   `json:"resource_type"`
	State         FreshnessState `json:"state"`
	ObservedAt    *time.Time     `json:"observed_at,omitempty"`
	LastIndexedAt *time.Time     `json:"last_indexed_at,omitempty"`
}

type Freshness struct {
	State      FreshnessState      `json:"state"`
	ObservedAt *time.Time          `json:"observed_at,omitempty"`
	Resources  []ResourceFreshness `json:"resources,omitempty"`
}

// Page contains resource projections and continuation metadata.
type Page struct {
	ResourceTypes  []ResourceType  `json:"resource_types"`
	Sort           Sort            `json:"sort"`
	PageSize       int             `json:"page_size"`
	ResultCount    int             `json:"result_count"`
	HasMore        bool            `json:"has_more"`
	Total          Total           `json:"total"`
	ResourceTotals []ResourceTotal `json:"resource_totals"`
	Results        []Result        `json:"results"`
	Facets         []FacetResult   `json:"facets"`
	NextCursor     string          `json:"next_cursor,omitempty"`
	Freshness      Freshness       `json:"freshness"`
}
