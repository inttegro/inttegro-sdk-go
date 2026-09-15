package paymentmethod

import (
	"encoding/json"
	"testing"
)

func TestPaymentMethodFingerprintJSON(t *testing.T) {
	tests := []struct {
		name        string
		jsonValue   string
		fingerprint string
	}{
		{
			name:        "fingerprinted method",
			jsonValue:   `{"id":"pm_1","fingerprint":"ifp_v1_app_customer"}`,
			fingerprint: "ifp_v1_app_customer",
		},
		{
			name:      "awaiting backfill",
			jsonValue: `{"id":"pm_2","fingerprint":""}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var method PaymentMethod
			if err := json.Unmarshal([]byte(tt.jsonValue), &method); err != nil {
				t.Fatalf("unmarshal payment method: %v", err)
			}
			if method.Fingerprint != tt.fingerprint {
				t.Fatalf("fingerprint = %q, want %q", method.Fingerprint, tt.fingerprint)
			}

			encoded, err := json.Marshal(method)
			if err != nil {
				t.Fatalf("marshal payment method: %v", err)
			}
			var projected map[string]any
			if err := json.Unmarshal(encoded, &projected); err != nil {
				t.Fatalf("unmarshal projected payment method: %v", err)
			}
			if projected["fingerprint"] != tt.fingerprint {
				t.Fatalf("marshaled fingerprint = %#v, want %q", projected["fingerprint"], tt.fingerprint)
			}
		})
	}
}
