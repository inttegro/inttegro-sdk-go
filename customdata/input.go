package customdata

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Input is open-ended custom data accepted by create and replace requests.
type Input struct {
	values map[string]json.RawMessage
}

// NewInput validates and copies an optional map into a new Input value.
func NewInput(values ...map[string]any) (*Input, error) {
	if len(values) > 1 {
		return nil, fmt.Errorf("inttegro: customdata.NewInput accepts at most one map")
	}
	input := &Input{values: make(map[string]json.RawMessage)}
	if len(values) == 0 || values[0] == nil {
		return input, nil
	}
	for key, value := range values[0] {
		if err := input.Set(key, value); err != nil {
			return nil, err
		}
	}
	return input, nil
}

func (i *Input) Set(key string, value any) error {
	if err := validateKey(key); err != nil {
		return err
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("inttegro: encode custom data %q: %w", key, err)
	}
	if i.values == nil {
		i.values = make(map[string]json.RawMessage)
	}
	previous, existed := i.values[key]
	i.values[key] = append(json.RawMessage(nil), encoded...)
	if err := validateSize(i.values); err != nil {
		if existed {
			i.values[key] = previous
		} else {
			delete(i.values, key)
		}
		return err
	}
	return nil
}

func (i *Input) Remove(key string) {
	if i != nil {
		delete(i.values, key)
	}
}

func (i *Input) Decode(key string, out any) (bool, error) {
	if i == nil {
		return false, nil
	}
	value, ok := i.values[key]
	if !ok {
		return false, nil
	}
	if out == nil {
		return true, errors.New("inttegro: custom data decode target is nil")
	}
	if err := json.Unmarshal(value, out); err != nil {
		return true, fmt.Errorf("inttegro: decode custom data %q: %w", key, err)
	}
	return true, nil
}

func (i *Input) Len() int {
	if i == nil {
		return 0
	}
	return len(i.values)
}

func (i Input) MarshalJSON() ([]byte, error) {
	if i.values == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(i.values)
}

func (i *Input) UnmarshalJSON(encoded []byte) error {
	if i == nil {
		return errors.New("inttegro: cannot decode custom data input into nil receiver")
	}
	if string(encoded) == "null" {
		i.values = nil
		return nil
	}
	var values map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &values); err != nil {
		return fmt.Errorf("inttegro: decode custom data input: %w", err)
	}
	if err := validateRawValues(values); err != nil {
		return err
	}
	i.values = cloneRawValues(values)
	return nil
}
