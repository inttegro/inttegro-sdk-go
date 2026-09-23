package customer

import (
	"time"

	"github.com/zebodotdev/inttegro-sdk-go/v10/customdata"
	"github.com/zebodotdev/inttegro-sdk-go/v10/money"
)

// CustomerData captures inline customer information for order creation.
//
// Use this when creating orders for new customers who don't have a customer ID yet.
// The API will create a customer record automatically and return it in the response.
//
// For existing customers, pass CustomerID instead to link the order to their record.
//
// Example:
//
//	customerData := &customer.Data{
//	    Name:        "Jane Doe",
//	    Email:       "jane@example.com",
//	    PhoneNumber: "+233244123456",
//	    Reference:   "customer_123_in_my_system",
//	}
type Data struct {
	// Name is the customer's full name (required).
	// Used for billing records and invoice generation.
	Name string `json:"name"`

	// Email is the customer's email address (optional but recommended).
	// Used for sending invoice links and payment notifications.
	// Validated as RFC 5322 email format.
	Email string `json:"email_address,omitempty"`

	// PhoneNumber is the customer's phone number with country code (required).
	// Used for SMS notifications and payment confirmations.
	// Example: "+233244123456"
	PhoneNumber string `json:"phone_number"`

	// Reference is your internal customer identifier (optional).
	// Use this to link Inttegro customer records to your system's users.
	// Maximum 255 characters. Must be unique across your customers.
	Reference string `json:"reference,omitempty"`

	CustomData *customdata.Input `json:"custom_data,omitempty"`
}

// Customer represents a customer record.
type Customer struct {
	Balance         map[string]BalanceValue `json:"balance,omitempty"`
	BillingAddress  *Address                `json:"billing_address,omitempty"`
	CreatedAt       *time.Time              `json:"created_at,omitempty"`
	CustomData      *customdata.Data        `json:"custom_data,omitempty"`
	Email           string                  `json:"email_address,omitempty"`
	Guest           bool                    `json:"guest,omitempty"`
	ID              string                  `json:"id,omitempty"`
	Name            string                  `json:"name,omitempty"`
	PhoneNumber     string                  `json:"phone_number,omitempty"`
	Reference       string                  `json:"reference,omitempty"`
	ShippingAddress *Address                `json:"shipping_address,omitempty"`
	Suffix          string                  `json:"suffix,omitempty"`
	Title           string                  `json:"title,omitempty"`
	UpdatedAt       *time.Time              `json:"updated_at,omitempty"`
}

// BalanceValue describes the customer's available balance in one currency.
type BalanceValue struct {
	AsOf      *time.Time   `json:"as_of,omitempty"`
	Available money.Amount `json:"available"`
}

// CustomersPage holds a page of customers.
type Page struct {
	Number    int        `json:"number,omitempty"`
	Size      int        `json:"size,omitempty"`
	Customers []Customer `json:"customers,omitempty"`
}
