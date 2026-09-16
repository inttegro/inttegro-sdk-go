package customdata

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Data is merchant-defined string data returned on an API resource.
type Data struct {
	values map[string]string
}

// New validates and copies values into a new Data value.
func New(values ...map[string]string) (*Data, error) {
	if len(values) > 1 {
		return nil, fmt.Errorf("inttegro: customdata.New accepts at most one map")
	}
	data := &Data{values: make(map[string]string)}
	if len(values) == 0 || values[0] == nil {
		return data, nil
	}
	for key, value := range values[0] {
		if err := data.Set(key, value); err != nil {
			return nil, err
		}
	}
	return data, nil
}

func (d *Data) Set(key, value string) error {
	if err := validateKey(key); err != nil {
		return err
	}
	if d.values == nil {
		d.values = make(map[string]string)
	}
	previous, existed := d.values[key]
	d.values[key] = value
	if err := validateSize(d.values); err != nil {
		if existed {
			d.values[key] = previous
		} else {
			delete(d.values, key)
		}
		return err
	}
	return nil
}

func (d *Data) Remove(key string) {
	if d != nil {
		delete(d.values, key)
	}
}

func (d *Data) Get(key string) (string, bool) {
	if d == nil {
		return "", false
	}
	value, ok := d.values[key]
	return value, ok
}

func (d *Data) Len() int {
	if d == nil {
		return 0
	}
	return len(d.values)
}

// Values returns a defensive copy of the stored values.
func (d *Data) Values() map[string]string {
	if d == nil {
		return nil
	}
	values := make(map[string]string, len(d.values))
	for key, value := range d.values {
		values[key] = value
	}
	return values
}

func (d *Data) Clone() *Data {
	clone, _ := New(d.Values())
	return clone
}

func (d Data) MarshalJSON() ([]byte, error) {
	if d.values == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(d.values)
}

func (d *Data) UnmarshalJSON(encoded []byte) error {
	if d == nil {
		return errors.New("inttegro: cannot decode custom data into nil receiver")
	}
	if string(encoded) == "null" {
		d.values = nil
		return nil
	}
	var values map[string]string
	if err := json.Unmarshal(encoded, &values); err != nil {
		return fmt.Errorf("inttegro: decode custom data: %w", err)
	}
	validated, err := New(values)
	if err != nil {
		return err
	}
	d.values = validated.values
	return nil
}
