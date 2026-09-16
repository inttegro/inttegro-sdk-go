package order

import "github.com/zebodotdev/inttegro-sdk-go/v9/customdata"

type Invoice struct {
	Number string        `json:"number,omitempty"`
	Format InvoiceFormat `json:"format"`
}

type InvoiceFormat struct {
	Web     DocumentFormat  `json:"web"`
	PDF     DocumentFormat  `json:"pdf"`
	Receipt *DocumentFormat `json:"receipt,omitempty"`
}

type DocumentFormat struct {
	URL string `json:"url"`
}

type InvoiceSettings struct {
	Number     string           `json:"number,omitempty"`
	Memo       string           `json:"memo,omitempty"`
	Footer     string           `json:"footer,omitempty"`
	CustomData *customdata.Data `json:"custom_data,omitempty"`
}
