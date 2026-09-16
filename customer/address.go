package customer

// Address represents a postal address used for billing and shipping.
type Address struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Line1       string `json:"line1"`
	Line2       string `json:"line2,omitempty"`
	City        string `json:"city,omitempty"`
	// Town is retained for backwards compatibility. New integrations should use City.
	Town   string `json:"town,omitempty"`
	Region string `json:"region,omitempty"`
	// District is retained for backwards compatibility with older payloads.
	District string `json:"district,omitempty"`
	Country  string `json:"country"`
	PostCode string `json:"post_code,omitempty"`
}
