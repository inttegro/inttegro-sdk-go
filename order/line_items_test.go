package order

import (
	"encoding/json"
	"testing"

	"github.com/zebodotdev/inttegro-sdk-go/v9/money"
	"github.com/zebodotdev/inttegro-sdk-go/v9/price"
)

func TestProductLineItemParamsMarshalCatalogProductWithExplicitPrice(t *testing.T) {
	value := ProductLineItemParams{
		ProductID: "prod_abc123xyz",
		Quantity:  2,
		Price: price.InlineParams{
			AmountParams: money.AmountParams{Currency: money.USD, Value: 4100},
		},
	}

	assertProductLineItemJSON(t, value, `{"product_id":"prod_abc123xyz","quantity":2,"price":{"currency":"usd","value":4100}}`)
}

func TestProductLineItemParamsMarshalCatalogProductWithSavedPrice(t *testing.T) {
	value := ProductLineItemParams{
		ProductID: "prod_abc123xyz",
		PriceID:   "pr_xyz789",
		Quantity:  2,
	}

	assertProductLineItemJSON(t, value, `{"product_id":"prod_abc123xyz","price_id":"pr_xyz789","quantity":2}`)
}

func assertProductLineItemJSON(t *testing.T, value ProductLineItemParams, want string) {
	t.Helper()

	got, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal product line item: %v", err)
	}
	if string(got) != want {
		t.Fatalf("marshaled product line item = %s, want %s", got, want)
	}
}
