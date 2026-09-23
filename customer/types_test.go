package customer

import (
	"encoding/json"
	"testing"

	"github.com/zebodotdev/inttegro-sdk-go/v10/customdata"
)

func TestCreateParamsUseTypedAddressesAndCustomData(t *testing.T) {
	data, err := customdata.NewInput(map[string]any{"segment": "vip"})
	if err != nil {
		t.Fatal(err)
	}
	params := CreateParams{
		Name:            "Ama Mensah",
		BillingAddress:  &Address{City: "Accra", Country: "gh"},
		CustomData:      data,
		ShippingAddress: &Address{City: "Kumasi", Country: "gh"},
	}
	encoded, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	var payload map[string]any
	if err := json.Unmarshal(encoded, &payload); err != nil {
		t.Fatal(err)
	}
	if payload["billing_address"].(map[string]any)["city"] != "Accra" {
		t.Fatalf("payload = %#v", payload)
	}
	if payload["custom_data"].(map[string]any)["segment"] != "vip" {
		t.Fatalf("payload = %#v", payload)
	}
}

func TestCustomerDecodesTypedAddressesAndCustomData(t *testing.T) {
	var customer Customer
	err := json.Unmarshal([]byte(`{
		"id":"cu_1",
		"name":"Ama Mensah",
		"billing_address":{"city":"Accra","country":"gh"},
		"shipping_address":{"city":"Kumasi","country":"gh"},
		"custom_data":{"segment":"vip"}
	}`), &customer)
	if err != nil {
		t.Fatal(err)
	}
	if customer.BillingAddress == nil || customer.BillingAddress.City != "Accra" {
		t.Fatalf("billing address = %#v", customer.BillingAddress)
	}
	if customer.CustomData == nil {
		t.Fatal("custom data was not decoded")
	}
	if segment, ok := customer.CustomData.Get("segment"); !ok || segment != "vip" {
		t.Fatalf("segment = %q, present=%v", segment, ok)
	}
}
