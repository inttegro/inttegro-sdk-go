package inttegro

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/inttegro/inttegro-sdk-go/v12/childrenpolicy"
	"github.com/inttegro/inttegro-sdk-go/v12/operatingpolicy"
)

func TestPolicyVersionLookupAndPage(t *testing.T) {
	type capturedRequest struct {
		path string
		body map[string]any
	}
	var captured []capturedRequest
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		captured = append(captured, capturedRequest{path: r.URL.Path, body: body})
		w.Header().Set("Content-Type", "application/json")
		version := 3
		if requested, ok := body["version"].(float64); ok {
			version = int(requested)
		}
		policy := map[string]any{
			"id": "cp_123", "version": version, "active_version": 3, "name": "Default",
			"capabilities": map[string]any{}, "created_at": "2026-11-05T10:00:00Z",
			"updated_at": "2026-11-05T10:00:00Z", "version_created_at": "2026-11-05T10:00:00Z",
		}
		if r.URL.Path == "/operating_policies/lookup" || r.URL.Path == "/operating_policies/page" {
			policy["id"] = "op_123"
			policy["name"] = "Retail"
			policy["children_policy_id"] = "cp_123"
			policy["children_policy_version"] = 3
		}
		if r.URL.Path == "/children_policies/page" || r.URL.Path == "/operating_policies/page" {
			_ = json.NewEncoder(w).Encode(map[string]any{"page": map[string]any{
				"number": 2, "size": 1, "count": 1, "total": 3, "has_more": true,
				"versions": []any{policy},
			}})
			return
		}
		field := "children_policy"
		if r.URL.Path == "/operating_policies/lookup" {
			field = "operating_policy"
		}
		_ = json.NewEncoder(w).Encode(map[string]any{field: policy})
	}))
	if client == nil {
		return
	}
	defer close()

	ctx := context.Background()
	latestChildren, err := client.ChildrenPolicies.Lookup(ctx)
	if err != nil {
		t.Fatal(err)
	}
	historicalChildren, err := client.ChildrenPolicies.Lookup(ctx, childrenpolicy.LookupParams{Version: 2})
	if err != nil {
		t.Fatal(err)
	}
	childrenPage, err := client.ChildrenPolicies.Page(ctx, childrenpolicy.PageParams{Number: 2, Size: 1})
	if err != nil {
		t.Fatal(err)
	}
	latestOperating, err := client.OperatingPolicies.Lookup(ctx, operatingpolicy.LookupParams{OperatingPolicyID: "op_123"})
	if err != nil {
		t.Fatal(err)
	}
	historicalOperating, err := client.OperatingPolicies.Lookup(ctx, operatingpolicy.LookupParams{OperatingPolicyID: "op_123", Version: 2})
	if err != nil {
		t.Fatal(err)
	}
	operatingPage, err := client.OperatingPolicies.Page(ctx, operatingpolicy.PageParams{OperatingPolicyID: "op_123", Number: 2, Size: 1})
	if err != nil {
		t.Fatal(err)
	}

	if latestChildren.Version != 3 || historicalChildren.Version != 2 || latestOperating.Version != 3 || historicalOperating.Version != 2 {
		t.Fatalf("decoded lookup versions = children %d/%d, operating %d/%d", latestChildren.Version, historicalChildren.Version, latestOperating.Version, historicalOperating.Version)
	}
	if childrenPage.Total != 3 || !childrenPage.HasMore || len(childrenPage.Versions) != 1 || childrenPage.Versions[0].Version != 3 {
		t.Fatalf("decoded children page = %#v", childrenPage)
	}
	if operatingPage.Total != 3 || !operatingPage.HasMore || len(operatingPage.Versions) != 1 || operatingPage.Versions[0].ID != "op_123" {
		t.Fatalf("decoded operating page = %#v", operatingPage)
	}

	want := []capturedRequest{
		{path: "/children_policies/lookup", body: map[string]any{}},
		{path: "/children_policies/lookup", body: map[string]any{"version": float64(2)}},
		{path: "/children_policies/page", body: map[string]any{"number": float64(2), "size": float64(1)}},
		{path: "/operating_policies/lookup", body: map[string]any{"operating_policy_id": "op_123"}},
		{path: "/operating_policies/lookup", body: map[string]any{"operating_policy_id": "op_123", "version": float64(2)}},
		{path: "/operating_policies/page", body: map[string]any{"operating_policy_id": "op_123", "number": float64(2), "size": float64(1)}},
	}
	if len(captured) != len(want) {
		t.Fatalf("captured requests = %#v", captured)
	}
	for i := range want {
		if captured[i].path != want[i].path {
			t.Fatalf("request %d path = %q, want %q", i, captured[i].path, want[i].path)
		}
		assertJSONMapEqual(t, captured[i].body, want[i].body)
	}
}
