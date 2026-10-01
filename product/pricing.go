package product

import (
	"github.com/zebodotdev/inttegro-sdk-go/v10/money"
)

// PriceType identifies the definition carried by an embedded product price.
type PriceType string

const (
	// PriceTypeFixedAmount identifies a fixed product price.
	PriceTypeFixedAmount            PriceType = "fixed_amount"
	// PriceTypeCustomerSelectedAmount identifies a customer-selected product price.
	PriceTypeCustomerSelectedAmount PriceType = "customer_selected_amount"
)

// SuggestedAmount is a convenient customer choice embedded in a product price.
type SuggestedAmount struct {
	ID          string `json:"id"`
	Value       int64  `json:"value"`
	Recommended bool   `json:"recommended,omitempty"`
}

// CustomerSelectedAmount is the selection policy embedded in a product price.
type CustomerSelectedAmount struct {
	Currency         money.Currency    `json:"currency"`
	Minimum          int64             `json:"minimum"`
	Maximum          *int64            `json:"maximum,omitempty"`
	SuggestedAmounts []SuggestedAmount `json:"suggested_amounts,omitempty"`
}

// ProductPriceSummary represents a product price listed with a product response.
type PriceSummary struct {
	ID                     string                  `json:"id"`
	Label                  string                  `json:"label,omitempty"`
	Type                   PriceType               `json:"type"`
	Nominal                money.Amount            `json:"nominal,omitempty"`
	FixedAmount            *money.Amount           `json:"fixed_amount,omitempty"`
	CustomerSelectedAmount *CustomerSelectedAmount `json:"customer_selected_amount,omitempty"`
	Active                 bool                    `json:"active"`
}
