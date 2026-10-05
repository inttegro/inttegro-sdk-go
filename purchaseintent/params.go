package purchaseintent

import (
	"encoding/json"
	"time"

	"github.com/inttegro/inttegro-sdk-go/v11/price"
)

type OriginalPriceParams struct {
	ID      string              `json:"id,omitempty"`
	Nominal *price.InlineParams `json:"nominal,omitempty"`
}

// BuyPageText contains merchant-authored copy for a hosted Buy page.
type BuyPageText struct {
	CheckoutSectionTitle string `json:"checkout_section_title,omitempty"`
	AmountFieldLabel     string `json:"amount_field_label,omitempty"`
	PrimaryActionLabel   string `json:"primary_action_label,omitempty"`
}

// BuyPagePresentation configures the hosted Buy page surface.
type BuyPagePresentation struct {
	Text *BuyPageText `json:"text,omitempty"`
}

// Presentation configures how a purchase intent is presented to customers.
type Presentation struct {
	BuyPage *BuyPagePresentation `json:"buy_page,omitempty"`
}

// TextValueUpdate distinguishes a replacement string from an explicit JSON
// null, which restores the product-aware default.
type TextValueUpdate struct {
	value *string
}

// SetText replaces one merchant-authored label.
func SetText(value string) *TextValueUpdate {
	return &TextValueUpdate{value: &value}
}

// ClearText restores the product-aware default for one label.
func ClearText() *TextValueUpdate {
	return &TextValueUpdate{}
}

// MarshalJSON implements json.Marshaler.
func (u TextValueUpdate) MarshalJSON() ([]byte, error) {
	if u.value == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*u.value)
}

// BuyPageTextUpdate is a sparse update to hosted Buy page copy. Leave a field
// nil to preserve it, use SetText to replace it, or ClearText to restore its
// product-aware default.
type BuyPageTextUpdate struct {
	CheckoutSectionTitle *TextValueUpdate `json:"checkout_section_title,omitempty"`
	AmountFieldLabel     *TextValueUpdate `json:"amount_field_label,omitempty"`
	PrimaryActionLabel   *TextValueUpdate `json:"primary_action_label,omitempty"`
}

type BuyPagePresentationUpdate struct {
	Text BuyPageTextUpdate `json:"text"`
}

type PresentationUpdate struct {
	BuyPage BuyPagePresentationUpdate `json:"buy_page"`
}

type CreateParams struct {
	Product      *ProductSelector `json:"product,omitempty"`
	ProductID    string           `json:"product_id,omitempty"`
	Price        *PriceSelector   `json:"price,omitempty"`
	PriceID      string           `json:"price_id,omitempty"`
	Quantity     Quantity         `json:"quantity"`
	Usage        *Usage           `json:"usage,omitempty"`
	ExpiresAt    *time.Time       `json:"expires_at,omitempty"`
	Presentation *Presentation    `json:"presentation,omitempty"`
}

type UpdateParams struct {
	ID           string              `json:"id"`
	Quantity     *Quantity           `json:"quantity,omitempty"`
	ExpiresAt    any                 `json:"expires_at,omitempty"`
	Reactivate   *bool               `json:"reactivate,omitempty"`
	Presentation *PresentationUpdate `json:"presentation,omitempty"`
}

type PageParams struct {
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}
