package inttegro

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FuzzRequestMetaIdempotency(f *testing.F) {
	for _, seed := range []string{
		`{"amount": 100}`,
		`{"idempotency_key":"legacy","request_meta":{"idempotency_key":"existing"}}`,
		`null`,
		`[]`,
		`not-json`,
	} {
		f.Add([]byte(seed), true)
		f.Add([]byte(seed), false)
	}

	f.Fuzz(func(t *testing.T, raw []byte, generate bool) {
		var original map[string]any
		decodeErr := json.Unmarshal(raw, &original)
		_, hasTopLevelKey := original["idempotency_key"]

		encoded, err := withRequestMetaIdempotency(raw, generate)
		if decodeErr == nil && original != nil && hasTopLevelKey {
			if err == nil {
				t.Fatal("top-level idempotency_key was accepted")
			}
			return
		}
		if err != nil {
			t.Fatalf("transform request metadata: %v", err)
		}
		if json.Valid(raw) && !json.Valid(encoded) {
			t.Fatalf("valid input produced invalid JSON: %q", encoded)
		}

		var payload map[string]any
		if err := json.Unmarshal(encoded, &payload); err != nil || payload == nil {
			return
		}
	})
}

func FuzzMutationPathClassification(f *testing.F) {
	for _, seed := range []string{
		"/orders/create",
		"/orders/lookup",
		"/financial_accounts/balances",
		"/files/contents",
		"/payouts/page",
		"/",
	} {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, path string) {
		if !strings.HasPrefix(path, "/") || strings.ContainsAny(path, "?#\x00\r\n") {
			return
		}
		relative := isIdempotentMutationPath(path)
		absolute := isIdempotentMutationPath("https://api.example.test" + path)
		if relative != absolute {
			t.Fatalf("relative and absolute path classification differ for %q", path)
		}
	})
}

func FuzzAPIErrorDecoding(f *testing.F) {
	for _, seed := range []string{
		`{"code":"invalid_request","message":"bad input"}`,
		`{"error":{"code":"forbidden","message":"not allowed"}}`,
		`{"error":null}`,
		`not-json`,
		``,
	} {
		f.Add([]byte(seed))
	}

	f.Fuzz(func(t *testing.T, body []byte) {
		if len(body) > 1<<20 {
			t.Skip()
		}
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(body)
		}))
		defer server.Close()

		client := NewClient("test-key", WithBaseURL(server.URL))
		err := client.do(context.Background(), http.MethodGet, "/fuzz", nil, nil)
		var apiErr *APIError
		if !errors.As(err, &apiErr) {
			t.Fatalf("expected APIError, got %T: %v", err, err)
		}
		if apiErr.StatusCode != http.StatusBadRequest {
			t.Fatalf("status code = %d", apiErr.StatusCode)
		}
	})
}
