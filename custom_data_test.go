package inttegro

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

func TestCustomDataCopiesAndControlsValues(t *testing.T) {
	source := map[string]string{"campaign": "launch"}
	data, err := NewCustomData(source)
	if err != nil {
		t.Fatal(err)
	}
	source["campaign"] = "mutated"
	if got, _ := data.Get("campaign"); got != "launch" {
		t.Fatalf("stored value = %q", got)
	}

	values := data.Values()
	values["campaign"] = "also mutated"
	if got, _ := data.Get("campaign"); got != "launch" {
		t.Fatalf("stored value after Values mutation = %q", got)
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != `{"campaign":"launch"}` {
		t.Fatalf("encoded = %s", encoded)
	}
}

func TestCustomDataRejectsInvalidLimitsWithoutMutation(t *testing.T) {
	data, _ := NewCustomData(map[string]string{"safe": "value"})
	if err := data.Set(strings.Repeat("k", MaxCustomDataKeyBytes+1), "value"); !errors.Is(err, ErrCustomDataKeyTooLong) {
		t.Fatalf("error = %v", err)
	}
	if err := data.Set("large", strings.Repeat("v", MaxCustomDataBytes)); !errors.Is(err, ErrCustomDataTooLarge) {
		t.Fatalf("error = %v", err)
	}
	if _, ok := data.Get("large"); ok {
		t.Fatal("rejected value was retained")
	}
}

func TestNewCustomDataRejectsMultipleSourceMaps(t *testing.T) {
	if _, err := NewCustomData(nil, map[string]string{"campaign": "launch"}); err == nil {
		t.Fatal("expected multiple source maps to be rejected")
	}
}

func TestCustomDataPatchDistinguishesSetAndUnset(t *testing.T) {
	patch := NewCustomDataPatch()
	if err := patch.Set("segment", "vip"); err != nil {
		t.Fatal(err)
	}
	if err := patch.Unset("old_note"); err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(patch)
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]*string
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded["segment"] == nil || *decoded["segment"] != "vip" {
		t.Fatalf("segment = %#v", decoded["segment"])
	}
	if value, ok := decoded["old_note"]; !ok || value != nil {
		t.Fatalf("old_note = %#v, present=%v", value, ok)
	}
}

func TestCustomDataInputPreservesStructuredJSON(t *testing.T) {
	input, err := NewCustomDataInput(map[string]any{
		"campaign": "launch",
		"audience": map[string]any{"tier": "vip"},
		"visits":   3,
	})
	if err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded["campaign"] != "launch" || decoded["visits"] != float64(3) {
		t.Fatalf("decoded input = %#v", decoded)
	}
}

func TestCustomDataPatchAcceptsStructuredValues(t *testing.T) {
	patch := NewCustomDataPatch()
	if err := patch.Set("preferences", map[string]any{"newsletter": true}); err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(patch)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(encoded), `{"preferences":{"newsletter":true}}`; got != want {
		t.Fatalf("encoded = %s, want %s", got, want)
	}
}
