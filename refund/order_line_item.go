package refund

import (
	"encoding/json"
	"fmt"
)

// OrderLineItemType identifies the kind of order line captured by a refund.
type OrderLineItemType string

const (
	OrderLineItemTypeProduct  OrderLineItemType = "product"
	OrderLineItemTypeFee      OrderLineItemType = "fee"
	OrderLineItemTypeShipping OrderLineItemType = "shipping"
)

// OrderLineItem is the immutable order-line snapshot attached to a refund.
// Exactly one detail object is present according to Type.
type OrderLineItem struct {
	ID       string                   `json:"id"`
	Type     OrderLineItemType        `json:"type"`
	Quantity int                      `json:"quantity,omitempty"`
	Product  *OrderLineItemProduct    `json:"product,omitempty"`
	Fee      *OrderLineItemAdjustment `json:"fee,omitempty"`
	Shipping *OrderLineItemAdjustment `json:"shipping,omitempty"`
}

// UnmarshalJSON rejects unsupported types and mismatched detail objects.
func (i *OrderLineItem) UnmarshalJSON(data []byte) error {
	var wire struct {
		ID       string            `json:"id"`
		Type     OrderLineItemType `json:"type"`
		Quantity *int              `json:"quantity"`
		Product  json.RawMessage   `json:"product"`
		Fee      json.RawMessage   `json:"fee"`
		Shipping json.RawMessage   `json:"shipping"`
	}
	if err := decodeSettlementJSON(data, &wire); err != nil {
		return fmt.Errorf("decode refund order line item: %w", err)
	}

	i.ID = wire.ID
	i.Type = wire.Type
	i.Quantity = 0
	i.Product = nil
	i.Fee = nil
	i.Shipping = nil
	switch wire.Type {
	case OrderLineItemTypeProduct:
		if wire.Quantity == nil || *wire.Quantity < 1 || missingJSONValue(wire.Product) ||
			len(wire.Fee) != 0 || len(wire.Shipping) != 0 {
			return fmt.Errorf("decode refund order line item: invalid product details")
		}
		var product OrderLineItemProduct
		if err := decodeSettlementJSON(wire.Product, &product); err != nil {
			return fmt.Errorf("decode refund order line item product: %w", err)
		}
		i.Quantity = *wire.Quantity
		i.Product = &product
	case OrderLineItemTypeFee:
		if wire.Quantity != nil || missingJSONValue(wire.Fee) || len(wire.Product) != 0 || len(wire.Shipping) != 0 {
			return fmt.Errorf("decode refund order line item: invalid fee details")
		}
		var fee OrderLineItemAdjustment
		if err := decodeSettlementJSON(wire.Fee, &fee); err != nil {
			return fmt.Errorf("decode refund order line item fee: %w", err)
		}
		i.Fee = &fee
	case OrderLineItemTypeShipping:
		if wire.Quantity != nil || missingJSONValue(wire.Shipping) || len(wire.Product) != 0 || len(wire.Fee) != 0 {
			return fmt.Errorf("decode refund order line item: invalid shipping details")
		}
		var shipping OrderLineItemAdjustment
		if err := decodeSettlementJSON(wire.Shipping, &shipping); err != nil {
			return fmt.Errorf("decode refund order line item shipping: %w", err)
		}
		i.Shipping = &shipping
	default:
		return fmt.Errorf("decode refund order line item: unsupported type %q", wire.Type)
	}
	return nil
}

// OrderLineItemProduct identifies the product captured by a product line.
// ID is absent when the order used an inline product.
type OrderLineItemProduct struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

// OrderLineItemAdjustment describes a fee or shipping line.
type OrderLineItemAdjustment struct {
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
}
