package product_test

import (
	"context"

	inttegro "github.com/zebodotdev/inttegro-sdk-go/v8"
	"github.com/zebodotdev/inttegro-sdk-go/v8/product"
)

func ExampleService_Create() {
	client := inttegro.NewClient("sk_test_example")
	_, _ = client.Products.Create(context.Background(), product.CreateParams{
		Type: product.TypeDigital,
		Name: "Monthly subscription",
	})
}
