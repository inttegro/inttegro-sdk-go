package inttegro

import (
	"encoding/json"
	"testing"

	"github.com/inttegro/inttegro-sdk-go/v10/money"
	"github.com/inttegro/inttegro-sdk-go/v10/price"
)

func TestPriceParamsEmbedsAmountOnTheWire(t *testing.T) {
	payload, err := json.Marshal(price.InlineParams{
		AmountParams: money.AmountParams{Currency: money.GHS, Value: 3005},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(payload), `{"currency":"ghs","value":3005}`; got != want {
		t.Fatalf("json.Marshal(PriceParams) = %s, want %s", got, want)
	}
}

func TestCatalogPriceRetainsReferencedProductID(t *testing.T) {
	var price price.Price
	if err := json.Unmarshal([]byte(`{"id":"pr_123","active":true,"nominal":{"currency":"ghs","value":3005},"product_id":"prod_123","created_at":"2026-09-02T12:00:00Z"}`), &price); err != nil {
		t.Fatal(err)
	}
	if got, want := price.ProductID, "prod_123"; got != want {
		t.Fatalf("CatalogPrice.ProductID = %q, want %q", got, want)
	}
}

func TestCustomerSelectedPriceParamsUseTaggedWireShape(t *testing.T) {
	maximum := int64(500000)
	payload, err := json.Marshal(price.CreateParams{
		ProductID: "prod_donation",
		Type:      price.TypeCustomerSelectedAmount,
		CustomerSelectedAmount: &price.CustomerSelectedAmountParams{
			Currency: money.GHS,
			Minimum:  10000,
			Maximum:  &maximum,
			SuggestedAmounts: []price.SuggestedAmountParams{
				{ID: "supporter", Value: 50000},
				{ID: "champion", Value: 100000, Recommended: true},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	want := `{"product_id":"prod_donation","type":"customer_selected_amount","customer_selected_amount":{"currency":"ghs","minimum":10000,"maximum":500000,"suggested_amounts":[{"id":"supporter","value":50000},{"id":"champion","value":100000,"recommended":true}]}}`
	if string(payload) != want {
		t.Fatalf("json.Marshal(CreateParams) = %s, want %s", payload, want)
	}
}

func TestFixedPriceParamsRequireTaggedWireShape(t *testing.T) {
	payload, err := json.Marshal(price.CreateParams{
		Type:        price.TypeFixedAmount,
		FixedAmount: &money.AmountParams{Currency: money.GHS, Value: 3005},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(payload), `{"type":"fixed_amount","fixed_amount":{"currency":"ghs","value":3005}}`; got != want {
		t.Fatalf("json.Marshal(CreateParams) = %s, want %s", got, want)
	}

	if _, err := json.Marshal(price.CreateParams{}); err == nil {
		t.Fatal("json.Marshal(CreateParams{}) succeeded, want a missing definition error")
	}

	productPayload, err := json.Marshal(price.AddToProductParams{
		ProductID:   "prod_123",
		Type:        price.TypeFixedAmount,
		FixedAmount: &money.AmountParams{Currency: money.GHS, Value: 3005},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(productPayload), `{"product_id":"prod_123","type":"fixed_amount","fixed_amount":{"currency":"ghs","value":3005}}`; got != want {
		t.Fatalf("json.Marshal(AddToProductParams) = %s, want %s", got, want)
	}
}
