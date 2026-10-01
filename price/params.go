package price

import (
	"encoding/json"
	"fmt"

	"github.com/inttegro/inttegro-sdk-go/v10/money"
)

// Type identifies the definition carried by a catalog price.
type Type string

const (
	// TypeFixedAmount charges the catalog price's stored fixed amount.
	TypeFixedAmount Type = "fixed_amount"
	// TypeCustomerSelectedAmount validates an amount selected when an order is created.
	TypeCustomerSelectedAmount Type = "customer_selected_amount"
)

// SuggestedAmountParams is a convenient customer choice, not an allow-list entry.
type SuggestedAmountParams struct {
	ID          string `json:"id"`
	Value       int64  `json:"value"`
	Recommended bool   `json:"recommended,omitempty"`
}

// CustomerSelectedAmountParams defines the accepted currency and range for a
// customer-selected catalog price.
type CustomerSelectedAmountParams struct {
	Currency         money.Currency          `json:"currency"`
	Minimum          int64                   `json:"minimum"`
	Maximum          *int64                  `json:"maximum,omitempty"`
	SuggestedAmounts []SuggestedAmountParams `json:"suggested_amounts,omitempty"`
}

type PageParams struct {
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	ProductID  string `json:"product_id,omitempty"`
}

type ActionParams struct {
	PriceID string `json:"price_id"`
}

// AddToProductParams creates a new price for an existing product.
type AddToProductParams struct {
	ProductID              string                        `json:"product_id"`
	Label                  string                        `json:"label,omitempty"`
	About                  string                        `json:"about,omitempty"`
	Type                   Type                          `json:"-"`
	FixedAmount            *money.AmountParams           `json:"-"`
	CustomerSelectedAmount *CustomerSelectedAmountParams `json:"-"`
	SetAsDefault           bool                          `json:"set_as_default,omitempty"`
}

func (p AddToProductParams) MarshalJSON() ([]byte, error) {
	if p.ProductID == "" {
		return nil, fmt.Errorf("price: product_id is required")
	}
	definition, err := definitionWire(p.Type, p.FixedAmount, p.CustomerSelectedAmount)
	if err != nil {
		return nil, err
	}
	return json.Marshal(struct {
		ProductID              string                        `json:"product_id"`
		Label                  string                        `json:"label,omitempty"`
		About                  string                        `json:"about,omitempty"`
		Type                   Type                          `json:"type,omitempty"`
		FixedAmount            *money.AmountParams           `json:"fixed_amount,omitempty"`
		CustomerSelectedAmount *CustomerSelectedAmountParams `json:"customer_selected_amount,omitempty"`
		SetAsDefault           bool                          `json:"set_as_default,omitempty"`
	}{
		ProductID:              p.ProductID,
		Label:                  p.Label,
		About:                  p.About,
		Type:                   definition.priceType,
		FixedAmount:            definition.fixedAmount,
		CustomerSelectedAmount: definition.customerSelectedAmount,
		SetAsDefault:           p.SetAsDefault,
	})
}

// PriceParams is an inline price supplied in a request. It embeds the amount
// fields because the API's price shape is {currency, value}, not
// {amount: {currency, value}}.
type InlineParams struct {
	money.AmountParams
}

// CatalogPriceParams creates a catalog price.
type CreateParams struct {
	ProductID              string                        `json:"product_id,omitempty"`
	Label                  string                        `json:"label,omitempty"`
	About                  string                        `json:"about,omitempty"`
	Type                   Type                          `json:"-"`
	FixedAmount            *money.AmountParams           `json:"-"`
	CustomerSelectedAmount *CustomerSelectedAmountParams `json:"-"`
}

func (p CreateParams) MarshalJSON() ([]byte, error) {
	definition, err := definitionWire(p.Type, p.FixedAmount, p.CustomerSelectedAmount)
	if err != nil {
		return nil, err
	}
	if p.Type == TypeCustomerSelectedAmount && p.ProductID == "" {
		return nil, fmt.Errorf("price: customer-selected prices require product_id")
	}
	return json.Marshal(struct {
		ProductID              string                        `json:"product_id,omitempty"`
		Label                  string                        `json:"label,omitempty"`
		About                  string                        `json:"about,omitempty"`
		Type                   Type                          `json:"type,omitempty"`
		FixedAmount            *money.AmountParams           `json:"fixed_amount,omitempty"`
		CustomerSelectedAmount *CustomerSelectedAmountParams `json:"customer_selected_amount,omitempty"`
	}{
		ProductID:              p.ProductID,
		Label:                  p.Label,
		About:                  p.About,
		Type:                   definition.priceType,
		FixedAmount:            definition.fixedAmount,
		CustomerSelectedAmount: definition.customerSelectedAmount,
	})
}

type definitionJSON struct {
	priceType              Type
	fixedAmount            *money.AmountParams
	customerSelectedAmount *CustomerSelectedAmountParams
}

func definitionWire(priceType Type, fixed *money.AmountParams, selected *CustomerSelectedAmountParams) (definitionJSON, error) {
	if priceType == TypeFixedAmount && fixed != nil && selected == nil {
		return definitionJSON{priceType: priceType, fixedAmount: fixed}, nil
	}
	if priceType == TypeCustomerSelectedAmount && selected != nil && fixed == nil {
		return definitionJSON{priceType: priceType, customerSelectedAmount: selected}, nil
	}
	return definitionJSON{}, fmt.Errorf("price: provide exactly one valid price definition")
}

// LookupPriceParams looks up a price by ID.
type LookupParams struct {
	PriceID string `json:"price_id"`
}

// UpdatePriceParams updates price metadata.
type UpdateParams struct {
	PriceID   string `json:"price_id"`
	ProductID string `json:"product_id,omitempty"`
	Label     string `json:"label,omitempty"`
	About     string `json:"about,omitempty"`
}
