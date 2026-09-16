package customdata

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Patch records custom-data merge operations. Unset emits JSON null, while
// Remove only discards a pending local change.
type Patch struct {
	changes map[string]json.RawMessage
}

func NewPatch() *Patch {
	return &Patch{changes: make(map[string]json.RawMessage)}
}

func (p *Patch) Set(key string, value any) error {
	if value == nil {
		return errors.New("inttegro: use CustomDataPatch.Unset to remove a value")
	}
	if err := validateKey(key); err != nil {
		return err
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("inttegro: encode custom data patch %q: %w", key, err)
	}
	return p.change(key, encoded)
}

func (p *Patch) Unset(key string) error {
	if err := validateKey(key); err != nil {
		return err
	}
	return p.change(key, []byte("null"))
}

func (p *Patch) change(key string, encoded []byte) error {
	if p.changes == nil {
		p.changes = make(map[string]json.RawMessage)
	}
	previous, existed := p.changes[key]
	p.changes[key] = append(json.RawMessage(nil), encoded...)
	if err := validateSize(p.changes); err != nil {
		if existed {
			p.changes[key] = previous
		} else {
			delete(p.changes, key)
		}
		return err
	}
	return nil
}

func (p *Patch) Remove(key string) {
	if p != nil {
		delete(p.changes, key)
	}
}

func (p *Patch) Len() int {
	if p == nil {
		return 0
	}
	return len(p.changes)
}

func (p Patch) MarshalJSON() ([]byte, error) {
	if p.changes == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(p.changes)
}

func (p *Patch) UnmarshalJSON(encoded []byte) error {
	if p == nil {
		return errors.New("inttegro: cannot decode custom data patch into nil receiver")
	}
	var changes map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &changes); err != nil {
		return fmt.Errorf("inttegro: decode custom data patch: %w", err)
	}
	if err := validateRawValues(changes); err != nil {
		return err
	}
	p.changes = cloneRawValues(changes)
	return nil
}
