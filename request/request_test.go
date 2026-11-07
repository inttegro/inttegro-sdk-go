package request

import (
	"encoding/json"
	"testing"
)

func TestMetaSerializesResponseProjection(t *testing.T) {
	meta := Meta{
		Response: &Response{
			Projection: Projection{
				"order":   Pick("id", "status", "customer.name"),
				"payment": Omit("latest_error"),
			},
		},
	}

	encoded, err := json.Marshal(meta)
	if err != nil {
		t.Fatalf("marshal request metadata: %v", err)
	}

	want := `{"response":{"projection":{"order":{"pick":["id","status","customer.name"]},"payment":{"omit":["latest_error"]}}}}`
	if string(encoded) != want {
		t.Fatalf("request metadata = %s, want %s", encoded, want)
	}
}
