package inttegro

import (
	"encoding/json"
	"testing"

	"github.com/inttegro/inttegro-sdk-go/v12/childrenpolicy"
	"github.com/inttegro/inttegro-sdk-go/v12/operatingpolicy"
)

func TestPolicyAgingSpecsExposePublicID(t *testing.T) {
	childrenJSON := []byte(`{"public_id":"bt_aging_spec_immediate","t_plus":"0h"}`)
	operatingJSON := []byte(`{"public_id":"bt_aging_spec_168_hourly","t_plus":"168h"}`)

	var children childrenpolicy.BTAgingSpec
	if err := json.Unmarshal(childrenJSON, &children); err != nil {
		t.Fatal(err)
	}
	var operating operatingpolicy.BTAgingSpec
	if err := json.Unmarshal(operatingJSON, &operating); err != nil {
		t.Fatal(err)
	}

	if children.PublicID != "bt_aging_spec_immediate" || operating.PublicID != "bt_aging_spec_168_hourly" {
		t.Fatalf("decoded public ids = %q/%q", children.PublicID, operating.PublicID)
	}
}
