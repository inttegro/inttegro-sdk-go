package chime

import (
	"time"

	"github.com/inttegro/inttegro-sdk-go/v14/customdata"
)

type Page struct {
	Number int     `json:"number,omitempty"`
	Size   int     `json:"size,omitempty"`
	Chimes []Chime `json:"chimes,omitempty"`
}

// Chime represents a notification message.
//
// Chimes track delivery status, transmission details, and any errors
// that occurred during sending.
type Chime struct {
	// ID is the unique chime identifier (read-only).
	// Starts with "chm_". Example: "chm_abc123def456"
	ID string `json:"id,omitempty"`

	// CreatedAt is when the chime was created (ISO 8601, read-only).
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// FullMessage is the message content that was sent.
	FullMessage string `json:"full_message,omitempty"`

	// Recipient contains the recipient details.
	Recipient *Recipient `json:"recipient,omitempty"`

	// SenderID is the sender name shown to the recipient.
	SenderID string `json:"sender_id,omitempty"`

	// Purpose is the chime category.
	Purpose string `json:"purpose,omitempty"`

	// CustomData contains attached custom data.
	CustomData *customdata.Data `json:"custom_data,omitempty"`

	// Transmission contains low-level transmission details.
	Transmission *Transmission `json:"transmission,omitempty"`
}

// Transmission describes the public delivery lifecycle for a chime.
type Transmission struct {
	Address       string     `json:"address"`
	CreatedAt     time.Time  `json:"created_at"`
	DeliveredAt   *time.Time `json:"delivered_at,omitempty"`
	FailedAt      *time.Time `json:"failed_at,omitempty"`
	InitializedAt time.Time  `json:"initialized_at"`
	Mechanism     string     `json:"mechanism"`
	SentAt        *time.Time `json:"sent_at,omitempty"`
	SentVia       string     `json:"sent_via,omitempty"`
	Status        string     `json:"status"`
	SuppressedAt  *time.Time `json:"suppressed_at,omitempty"`
}

// Settings contains the effective non-secret delivery defaults for Chimes.
type Settings struct {
	SMS   SMSSettings   `json:"sms"`
	Email EmailSettings `json:"email"`
}

// SMSSettings describes the platform SMS sender defaults and accepted sender length.
type SMSSettings struct {
	DefaultSender    string `json:"default_sender"`
	MinSenderCharLen int    `json:"min_sender_char_len"`
	MaxSenderCharLen int    `json:"max_sender_char_len"`
}

// EmailSettings describes the default and allowed From addresses for email Chimes.
type EmailSettings struct {
	DefaultFrom string   `json:"default_from"`
	AllowedFrom []string `json:"allowed_from"`
}
