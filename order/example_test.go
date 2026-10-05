package order_test

import (
	"context"

	inttegro "github.com/inttegro/inttegro-sdk-go/v11"
	"github.com/inttegro/inttegro-sdk-go/v11/money"
	"github.com/inttegro/inttegro-sdk-go/v11/order"
	"github.com/inttegro/inttegro-sdk-go/v11/price"
	"github.com/inttegro/inttegro-sdk-go/v11/product"
)

func ExampleService_Create() {
	client := inttegro.NewClient("sk_test_example")
	_, _ = client.Orders.Create(context.Background(), order.CreateParams{
		LineItems: []order.LineItemParams{{
			Type: order.LineItemTypeProduct,
			Product: &order.ProductLineItemParams{
				Type:     product.TypeDigital,
				Name:     "Monthly subscription",
				Quantity: 1,
				Price: price.InlineParams{AmountParams: money.AmountParams{
					Currency: money.GHS,
					Value:    5000,
				}},
			},
		}},
	})
}
