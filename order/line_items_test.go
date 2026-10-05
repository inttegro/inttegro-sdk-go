package order

import (
	"encoding/json"
	"testing"

	"github.com/inttegro/inttegro-sdk-go/v11/money"
	"github.com/inttegro/inttegro-sdk-go/v11/price"
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

func TestProductLineItemParamsMarshalCatalogProductWithCustomerSelectedPrice(t *testing.T) {
	value := ProductLineItemParams{
		ProductID: "prod_donation",
		CustomerSelectedPrice: &CustomerSelectedPriceParams{
			PriceID:        "pr_donation",
			SelectedAmount: money.AmountParams{Currency: money.GHS, Value: 75000},
		},
		Quantity: 1,
	}

	assertProductLineItemJSON(t, value, `{"product_id":"prod_donation","customer_selected_price":{"price_id":"pr_donation","selected_amount":{"currency":"ghs","value":75000}},"quantity":1}`)
}

func TestProductLineItemParamsRejectsMixedCustomerSelectedPrice(t *testing.T) {
	value := ProductLineItemParams{
		ProductID: "prod_donation",
		PriceID:   "pr_fixed",
		CustomerSelectedPrice: &CustomerSelectedPriceParams{
			PriceID:        "pr_donation",
			SelectedAmount: money.AmountParams{Currency: money.GHS, Value: 75000},
		},
		Quantity: 1,
	}

	if _, err := json.Marshal(value); err == nil {
		t.Fatal("json.Marshal() error = nil, want mixed price choice rejection")
	}
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
