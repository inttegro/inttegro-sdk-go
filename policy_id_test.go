package inttegro

import (
	"encoding/json"
	"testing"

	"github.com/inttegro/inttegro-sdk-go/v13/childrenpolicy"
	"github.com/inttegro/inttegro-sdk-go/v13/operatingpolicy"
)

func TestPolicyAgingSpecsExposeID(t *testing.T) {
	childrenJSON := []byte(`{"id":"bt_aging_spec_immediate","t_plus":"0h"}`)
	operatingJSON := []byte(`{"id":"bt_aging_spec_168_hourly","t_plus":"168h"}`)

	var children childrenpolicy.BTAgingSpec
	if err := json.Unmarshal(childrenJSON, &children); err != nil {
		t.Fatal(err)
	}
	var operating operatingpolicy.BTAgingSpec
	if err := json.Unmarshal(operatingJSON, &operating); err != nil {
		t.Fatal(err)
	}

	if children.ID != "bt_aging_spec_immediate" || operating.ID != "bt_aging_spec_168_hourly" {
		t.Fatalf("decoded ids = %q/%q", children.ID, operating.ID)
	}
}
