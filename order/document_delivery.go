package order

import "github.com/inttegro/inttegro-sdk-go/v11/invoice"

// DocumentDeliveryResult contains the updated order and its delivery result.
type DocumentDeliveryResult struct {
	Order    Order            `json:"order"`
	Delivery invoice.Delivery `json:"delivery"`
}
