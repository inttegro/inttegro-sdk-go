package price

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v10/money"
	"github.com/zebodotdev/inttegro-sdk-go/v10/product"
)

type Page struct {
	Number int     `json:"number,omitempty"`
	Size   int     `json:"size,omitempty"`
	Prices []Price `json:"prices,omitempty"`
}

// Price is an inline price returned by the API.
type Inline struct {
	money.Amount
}

// SuggestedAmount is a convenient customer choice returned with a price.
type SuggestedAmount struct {
	ID          string `json:"id"`
	Value       int64  `json:"value"`
	Recommended bool   `json:"recommended,omitempty"`
}

// CustomerSelectedAmount is the persisted selection policy returned with a
// customer-selected catalog price.
type CustomerSelectedAmount struct {
	Currency         money.Currency    `json:"currency"`
	Minimum          int64             `json:"minimum"`
	Maximum          *int64            `json:"maximum,omitempty"`
	SuggestedAmounts []SuggestedAmount `json:"suggested_amounts,omitempty"`
}

// Price represents a catalog price.
type Price struct {
	ID                     string                  `json:"id,omitempty"`
	Label                  string                  `json:"label,omitempty"`
	About                  string                  `json:"about,omitempty"`
	Active                 bool                    `json:"active"`
	Type                   Type                    `json:"type"`
	Nominal                *money.Amount           `json:"nominal,omitempty"`
	FixedAmount            *money.Amount           `json:"fixed_amount,omitempty"`
	CustomerSelectedAmount *CustomerSelectedAmount `json:"customer_selected_amount,omitempty"`
	ProductID              string                  `json:"product_id,omitempty"`
	Product                *product.Product        `json:"product,omitempty"`
	CreatedAt              *time.Time              `json:"created_at,omitempty"`
	UpdatedAt              *time.Time              `json:"updated_at,omitempty"`
	ArchivedAt             *time.Time              `json:"archived_at,omitempty"`
}
