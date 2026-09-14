package refund

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// SettlementType discriminates the destination captured for a refund.
type SettlementType string

const (
	SettlementTypeOffline       SettlementType = "offline"
	SettlementTypePaymentMethod SettlementType = "payment_method"
)

// Settlement is the immutable destination snapshot for a refund.
// PaymentMethod is present only when Type is SettlementTypePaymentMethod.
type Settlement struct {
	Type          SettlementType           `json:"type"`
	PaymentMethod *SettlementPaymentMethod `json:"payment_method,omitempty"`
}

// UnmarshalJSON rejects impossible discriminator and payload combinations.
func (s *Settlement) UnmarshalJSON(data []byte) error {
	var wire struct {
		Type          SettlementType  `json:"type"`
		PaymentMethod json.RawMessage `json:"payment_method"`
	}
	if err := decodeSettlementJSON(data, &wire); err != nil {
		return fmt.Errorf("decode refund settlement: %w", err)
	}

	s.Type = wire.Type
	s.PaymentMethod = nil
	switch wire.Type {
	case SettlementTypeOffline:
		if len(wire.PaymentMethod) != 0 {
			return fmt.Errorf("decode refund settlement: offline must omit payment_method")
		}
	case SettlementTypePaymentMethod:
		if missingJSONValue(wire.PaymentMethod) {
			return fmt.Errorf("decode refund settlement: payment_method is required")
		}
		var method SettlementPaymentMethod
		if err := json.Unmarshal(wire.PaymentMethod, &method); err != nil {
			return fmt.Errorf("decode refund settlement payment_method: %w", err)
		}
		s.PaymentMethod = &method
	default:
		return fmt.Errorf("decode refund settlement: unsupported type %q", wire.Type)
	}
	return nil
}

// SettlementPaymentMethod is a caller-safe snapshot of the original payment
// method. Exactly one rail-specific detail object is present according to Type.
type SettlementPaymentMethod struct {
	ID          string                      `json:"id"`
	Type        SettlementPaymentMethodType `json:"type"`
	MobileMoney *SettlementMobileMoney      `json:"mobile_money,omitempty"`
	BankAccount *SettlementBankAccount      `json:"bank_account,omitempty"`
}

// UnmarshalJSON rejects unsupported rails and mismatched detail objects.
func (m *SettlementPaymentMethod) UnmarshalJSON(data []byte) error {
	var wire struct {
		ID          string                      `json:"id"`
		Type        SettlementPaymentMethodType `json:"type"`
		MobileMoney json.RawMessage             `json:"mobile_money"`
		BankAccount json.RawMessage             `json:"bank_account"`
	}
	if err := decodeSettlementJSON(data, &wire); err != nil {
		return fmt.Errorf("decode refund settlement payment method: %w", err)
	}
	if wire.ID == "" || wire.ID != strings.TrimSpace(wire.ID) || !strings.HasPrefix(wire.ID, "pm_") {
		return fmt.Errorf("decode refund settlement payment method: invalid public id")
	}

	m.ID = wire.ID
	m.Type = wire.Type
	m.MobileMoney = nil
	m.BankAccount = nil
	switch wire.Type {
	case SettlementPaymentMethodTypeMobileMoney:
		if missingJSONValue(wire.MobileMoney) || len(wire.BankAccount) != 0 {
			return fmt.Errorf("decode refund settlement payment method: invalid mobile_money details")
		}
		var details SettlementMobileMoney
		if err := decodeSettlementJSON(wire.MobileMoney, &details); err != nil {
			return fmt.Errorf("decode refund settlement mobile_money: %w", err)
		}
		if !validSettlementMask(details.AccountNumber, details.Last4) || !knownSettlementNetwork(details.Network) {
			return fmt.Errorf("decode refund settlement mobile_money: invalid masked details")
		}
		m.MobileMoney = &details
	case SettlementPaymentMethodTypeBankAccount:
		if missingJSONValue(wire.BankAccount) || len(wire.MobileMoney) != 0 {
			return fmt.Errorf("decode refund settlement payment method: invalid bank_account details")
		}
		var details SettlementBankAccount
		if err := json.Unmarshal(wire.BankAccount, &details); err != nil {
			return fmt.Errorf("decode refund settlement bank_account: %w", err)
		}
		m.BankAccount = &details
	default:
		return fmt.Errorf("decode refund settlement payment method: unsupported type %q", wire.Type)
	}
	return nil
}

// SettlementPaymentMethodType identifies a supported refund destination rail.
type SettlementPaymentMethodType string

const (
	SettlementPaymentMethodTypeMobileMoney SettlementPaymentMethodType = "mobile_money"
	SettlementPaymentMethodTypeBankAccount SettlementPaymentMethodType = "bank_account"
)

// SettlementMobileMoney contains only masked mobile-money recognition data.
type SettlementMobileMoney struct {
	Network       string `json:"network"`
	AccountNumber string `json:"account_number"`
	Last4         string `json:"last4"`
}

// SettlementBankAccount identifies the supported bank-account subtype.
type SettlementBankAccount struct {
	Type             SettlementBankAccountType   `json:"type"`
	GhanaBankAccount *SettlementGhanaBankAccount `json:"ghana_bank_account,omitempty"`
}

// UnmarshalJSON rejects unsupported account subtypes and unmasked details.
func (a *SettlementBankAccount) UnmarshalJSON(data []byte) error {
	var wire struct {
		Type             SettlementBankAccountType `json:"type"`
		GhanaBankAccount json.RawMessage           `json:"ghana_bank_account"`
	}
	if err := decodeSettlementJSON(data, &wire); err != nil {
		return fmt.Errorf("decode refund settlement bank account: %w", err)
	}
	if wire.Type != SettlementBankAccountTypeGhana || missingJSONValue(wire.GhanaBankAccount) {
		return fmt.Errorf("decode refund settlement bank account: unsupported type %q", wire.Type)
	}

	var details SettlementGhanaBankAccount
	if err := decodeSettlementJSON(wire.GhanaBankAccount, &details); err != nil {
		return fmt.Errorf("decode refund settlement ghana_bank_account: %w", err)
	}
	if !validSettlementMask(details.AccountNumber, details.Last4) {
		return fmt.Errorf("decode refund settlement ghana_bank_account: invalid masked details")
	}
	a.Type = wire.Type
	a.GhanaBankAccount = &details
	return nil
}

// SettlementBankAccountType identifies a supported bank-account snapshot.
type SettlementBankAccountType string

const SettlementBankAccountTypeGhana SettlementBankAccountType = "ghana_bank_account"

// SettlementGhanaBankAccount contains only masked account recognition data.
type SettlementGhanaBankAccount struct {
	AccountNumber string `json:"account_number"`
	Last4         string `json:"last4"`
}

func decodeSettlementJSON(data []byte, destination any) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(destination); err != nil {
		return err
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return fmt.Errorf("expected one JSON value")
	}
	return nil
}

func missingJSONValue(value json.RawMessage) bool {
	return len(value) == 0 || bytes.Equal(bytes.TrimSpace(value), []byte("null"))
}

func validSettlementMask(accountNumber, last4 string) bool {
	if len(accountNumber) != 8 || len(last4) != 4 || accountNumber[:4] != "****" {
		return false
	}
	for index := range last4 {
		if last4[index] < '0' || last4[index] > '9' {
			return false
		}
	}
	return accountNumber[4:] == last4
}

func knownSettlementNetwork(network string) bool {
	switch network {
	case "airtel", "mtn", "telecel", "vodafone":
		return true
	default:
		return false
	}
}
