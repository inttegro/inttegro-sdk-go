// Package customdata provides the semantic custom-data collections used by
// Inttegro request and response models.
package customdata

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	MaxKeyBytes = 256
	MaxBytes    = 25 * 1024
)

var (
	ErrKeyTooLong = errors.New("inttegro: custom data key exceeds 256 bytes")
	ErrTooLarge   = errors.New("inttegro: custom data exceeds 25 KiB")
)

func validateKey(key string) error {
	if len([]byte(key)) > MaxKeyBytes {
		return fmt.Errorf("%w: %q", ErrKeyTooLong, key)
	}
	return nil
}

func validateRawValues(values map[string]json.RawMessage) error {
	for key, value := range values {
		if err := validateKey(key); err != nil {
			return err
		}
		if !json.Valid(value) {
			return fmt.Errorf("inttegro: invalid JSON custom data value for %q", key)
		}
	}
	return validateSize(values)
}

func validateSize(values any) error {
	encoded, err := json.Marshal(values)
	if err != nil {
		return fmt.Errorf("inttegro: encode custom data: %w", err)
	}
	if len(encoded) > MaxBytes {
		return ErrTooLarge
	}
	return nil
}

func cloneRawValues(values map[string]json.RawMessage) map[string]json.RawMessage {
	clone := make(map[string]json.RawMessage, len(values))
	for key, value := range values {
		clone[key] = append(json.RawMessage(nil), value...)
	}
	return clone
}
