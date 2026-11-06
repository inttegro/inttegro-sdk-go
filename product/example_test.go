package product_test

import (
	"context"

	inttegro "github.com/inttegro/inttegro-sdk-go/v13"
	"github.com/inttegro/inttegro-sdk-go/v13/product"
)

func ExampleService_Create() {
	client := inttegro.NewClient("sk_test_example")
	_, _ = client.Products.Create(context.Background(), product.CreateParams{
		Type: product.TypeDigital,
		Name: "Monthly subscription",
	})
}
