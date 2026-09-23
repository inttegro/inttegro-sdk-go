package order

import "github.com/zebodotdev/inttegro-sdk-go/v10/invoice"

// DocumentDeliveryResult contains the updated order and its delivery result.
type DocumentDeliveryResult struct {
	Order    Order            `json:"order"`
	Delivery invoice.Delivery `json:"delivery"`
}
