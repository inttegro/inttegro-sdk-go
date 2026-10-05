package purchaseintent_test

import (
	"fmt"

	"github.com/inttegro/inttegro-sdk-go/v11/purchaseintent"
)

func ExampleStatus() {
	intent := purchaseintent.PurchaseIntent{Status: purchaseintent.StatusActive}
	fmt.Println(intent.Status)
	// Output: active
}
