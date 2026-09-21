package inttegro

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/zebodotdev/inttegro-sdk-go/v9/search"
)

func TestResourceSearchEndpointsUseTypedContract(t *testing.T) {
	var paths []string
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode search request: %v", err)
		}
		if _, ok := body["request_meta"]; ok {
			t.Fatalf("search request must not include mutation metadata: %#v", body)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"search": map[string]any{
			"resource_types":  []string{"order"},
			"sort":            map[string]any{"field": "relevance", "direction": "desc"},
			"page_size":       20,
			"result_count":    0,
			"has_more":        false,
			"total":           map[string]any{"value": 0, "relation": "exact"},
			"resource_totals": []any{},
			"results":         []any{},
			"facets":          []any{},
			"freshness":       map[string]any{"state": "current"},
		}})
	}))
	if client == nil {
		return
	}
	defer close()

	ctx := context.Background()
	params := search.Params{Text: "Ama Mensah"}
	if _, err := client.Customers.Search(ctx, params); err != nil {
		t.Fatal(err)
	}
	if _, err := client.FinancialAccounts.Search(ctx, params); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Orders.Search(ctx, params); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Payouts.Search(ctx, params); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Products.Search(ctx, params); err != nil {
		t.Fatal(err)
	}

	want := []string{
		"/customers/search",
		"/financial_accounts/search",
		"/orders/search",
		"/payouts/search",
		"/products/search",
	}
	if len(paths) != len(want) {
		t.Fatalf("paths = %v, want %v", paths, want)
	}
	for i := range want {
		if paths[i] != want[i] {
			t.Fatalf("path %d = %q, want %q", i, paths[i], want[i])
		}
	}
}
