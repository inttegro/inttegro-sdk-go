package purchaseintent

// Bundle is the resolved current catalog bundle for a bundle-backed Buy link.
type Bundle struct {
	ID        string        `json:"id"`
	Name      string        `json:"name,omitempty"`
	About     string        `json:"about,omitempty"`
	Reference string        `json:"reference,omitempty"`
	Media     *ProductMedia `json:"media,omitempty"`
	// Price is the current catalog reference; PurchaseIntent.Price is the offer.
	Price      *Price            `json:"price,omitempty"`
	Components []BundleComponent `json:"components,omitempty"`
}

// BundleComponent is the customer-visible component snapshot resolved for a
// bundle Buy link.
type BundleComponent struct {
	ProductID string `json:"product_id,omitempty"`
	PriceID   string `json:"price_id,omitempty"`
	Reference string `json:"reference,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
	About     string `json:"about,omitempty"`
	Category  string `json:"category,omitempty"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Type      string `json:"type,omitempty"`
}
