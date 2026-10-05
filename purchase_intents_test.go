package inttegro

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/inttegro/inttegro-sdk-go/v11/purchaseintent"
)

func TestPurchaseIntentLookupReturnsTypedResource(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/purchase_intents/lookup" {
			t.Fatalf("expected /purchase_intents/lookup, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"purchase_intent": map[string]any{
				"allow_variants": false,
				"created_at":     "2026-09-09T12:00:00Z",
				"id":             "sale_123",
				"merchant":       map[string]any{"organization_name": "Tea House Ltd"},
				"presentation": map[string]any{
					"buy_page": map[string]any{
						"text": map[string]any{"checkout_section_title": "Support this cause"},
					},
				},
				"product": map[string]any{
					"active":     true,
					"created_at": "2026-09-09T11:00:00Z",
					"dimensions": map[string]any{"digital": map[string]any{"bytes": 1024}},
					"id":         "prod_123",
					"name":       "Tea guide",
					"type":       "digital",
				},
				"quantity": map[string]any{"min": 1},
				"status":   "active",
				"usage": map[string]any{
					"order":      map[string]any{"created_at": "2026-09-09T12:02:00Z", "id": "or_123"},
					"single_use": true,
				},
			},
		})
	}))
	defer srv.Close()

	client := NewClient("sk_test", WithBaseURL(srv.URL))
	intent, err := client.PurchaseIntents.Lookup(context.Background(), "sale_123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := intent.Merchant.OrganizationName; got != "Tea House Ltd" {
		t.Fatalf("merchant organization name = %q", got)
	}
	if got := intent.Product.Dimensions.Digital.Bytes; got != 1024 {
		t.Fatalf("product byte size = %v", got)
	}
	if got := intent.Usage.Order.ID; got != "or_123" {
		t.Fatalf("usage order ID = %q", got)
	}
	if got := intent.Presentation.BuyPage.Text.CheckoutSectionTitle; got != "Support this cause" {
		t.Fatalf("checkout section title = %q", got)
	}
}

func TestPurchaseIntentPresentationSerializesCreateAndUpdate(t *testing.T) {
	createBody, err := json.Marshal(purchaseintent.CreateParams{
		Quantity: purchaseintent.Quantity{Min: 1},
		Presentation: &purchaseintent.Presentation{
			BuyPage: &purchaseintent.BuyPagePresentation{
				Text: &purchaseintent.BuyPageText{AmountFieldLabel: "Your contribution"},
			},
		},
	})
	if err != nil {
		t.Fatalf("marshal create params: %v", err)
	}
	if !strings.Contains(string(createBody), `"amount_field_label":"Your contribution"`) {
		t.Fatalf("create presentation missing from %s", createBody)
	}

	updateBody, err := json.Marshal(purchaseintent.UpdateParams{
		ID: "sale_123",
		Presentation: &purchaseintent.PresentationUpdate{
			BuyPage: purchaseintent.BuyPagePresentationUpdate{
				Text: purchaseintent.BuyPageTextUpdate{
					CheckoutSectionTitle: purchaseintent.SetText("Contribute now"),
					AmountFieldLabel:     purchaseintent.ClearText(),
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("marshal update params: %v", err)
	}
	if !strings.Contains(string(updateBody), `"checkout_section_title":"Contribute now"`) {
		t.Fatalf("update presentation missing from %s", updateBody)
	}
	if !strings.Contains(string(updateBody), `"amount_field_label":null`) {
		t.Fatalf("update presentation reset missing from %s", updateBody)
	}
	if strings.Contains(string(updateBody), `"primary_action_label"`) {
		t.Fatalf("omitted update field was serialized in %s", updateBody)
	}
}
