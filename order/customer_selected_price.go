package order

import "github.com/inttegro/inttegro-sdk-go/v10/money"

// CustomerSelectedPriceParams couples a saved customer-selected catalog price
// with the concrete unit amount chosen for an order.
type CustomerSelectedPriceParams struct {
	PriceID        string             `json:"price_id"`
	SelectedAmount money.AmountParams `json:"selected_amount"`
}
