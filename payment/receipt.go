package payment

import "time"

// Receipt is the receipt issued for a paid order payment.
type Receipt struct {
	Number          string                  `json:"number,omitempty"`
	Format          ReceiptFormat           `json:"format"`
	DeliverySummary *ReceiptDeliverySummary `json:"delivery_summary,omitempty"`
}

type ReceiptFormat struct {
	Web ReceiptDocumentFormat `json:"web"`
	PDF ReceiptDocumentFormat `json:"pdf"`
}

type ReceiptDocumentFormat struct {
	URL string `json:"url"`
}

type ReceiptDeliverySummary struct {
	Channels     map[string]ReceiptDeliveryChannelSummary `json:"channels"`
	TotalCount   int                                      `json:"total_count"`
	LatestSentAt *time.Time                               `json:"latest_sent_at,omitempty"`
}

type ReceiptDeliveryChannelSummary struct {
	Count        int        `json:"count"`
	LatestSentAt *time.Time `json:"latest_sent_at,omitempty"`
}
