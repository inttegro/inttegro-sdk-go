package refund

import (
	"encoding/json"
	"testing"
)

func TestSettlementUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var settlement Settlement
	err := json.Unmarshal([]byte(`{
		"type":"payment_method",
		"payment_method":{
			"id":"pm_123",
			"type":"bank_account",
			"bank_account":{
				"type":"ghana_bank_account",
				"ghana_bank_account":{"account_number":"****1234","last4":"1234"}
			}
		}
	}`), &settlement)
	if err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got := settlement.PaymentMethod.BankAccount.GhanaBankAccount.AccountNumber; got != "****1234" {
		t.Fatalf("AccountNumber = %q, want %q", got, "****1234")
	}
}

func TestSettlementUnmarshalJSONRejectsInvalidShapes(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"unknown settlement":        `{"type":"provider"}`,
		"offline with method":       `{"type":"offline","payment_method":{"id":"pm_123"}}`,
		"payment method missing":    `{"type":"payment_method"}`,
		"unsupported method":        `{"type":"payment_method","payment_method":{"id":"pm_123","type":"card"}}`,
		"unmasked mobile money":     `{"type":"payment_method","payment_method":{"id":"pm_123","type":"mobile_money","mobile_money":{"network":"mtn","account_number":"0244001234","last4":"1234"}}}`,
		"mismatched bank last four": `{"type":"payment_method","payment_method":{"id":"pm_123","type":"bank_account","bank_account":{"type":"ghana_bank_account","ghana_bank_account":{"account_number":"****1234","last4":"9999"}}}}`,
	}

	for name, value := range tests {
		name, value := name, value
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var settlement Settlement
			if err := json.Unmarshal([]byte(value), &settlement); err == nil {
				t.Fatal("Unmarshal() error = nil, want invalid settlement error")
			}
		})
	}
}
