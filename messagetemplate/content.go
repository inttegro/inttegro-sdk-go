package messagetemplate

type SMSContent struct {
	MessageTemplate string `json:"message_template"`
}

type Mailbox struct {
	Address string `json:"address"`
	Name    string `json:"name,omitempty"`
}

type EmailContent struct {
	From    *Mailbox          `json:"from,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	HTML    string            `json:"html"`
	ReplyTo *Mailbox          `json:"reply_to,omitempty"`
	Subject string            `json:"subject"`
}

type RenderedContent struct {
	Channel     Channel        `json:"channel"`
	Attachments []string       `json:"attachments,omitempty"`
	Email       *RenderedEmail `json:"email,omitempty"`
	SMS         *RenderedSMS   `json:"sms,omitempty"`
}

// RenderedSMS contains the final message sent through the SMS channel.
type RenderedSMS struct {
	FullMessage string `json:"full_message"`
}

// RenderedEmail contains the final email headers and bodies.
type RenderedEmail struct {
	Subject string            `json:"subject"`
	Text    string            `json:"text"`
	HTML    string            `json:"html,omitempty"`
	From    *Mailbox          `json:"from,omitempty"`
	ReplyTo *Mailbox          `json:"reply_to,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Safety  *SafetyResult     `json:"safety,omitempty"`
}

// SafetyResult describes the final content scan applied to rendered email.
type SafetyResult struct {
	ContentHash     string              `json:"content_hash"`
	Links           []ScannedLink       `json:"links,omitempty"`
	NormalizedText  string              `json:"normalized_text"`
	QuarantineNotes string              `json:"quarantine_notes,omitempty"`
	ReasonCodes     []string            `json:"reason_codes,omitempty"`
	SanitizedHTML   string              `json:"sanitized_html,omitempty"`
	Scanner         string              `json:"scanner"`
	Status          ContentSafetyStatus `json:"status"`
}

// ScannedLink contains the safety decision for one link in rendered email.
type ScannedLink struct {
	Host   string `json:"host,omitempty"`
	Raw    string `json:"raw"`
	Reason string `json:"reason,omitempty"`
	Scheme string `json:"scheme"`
	Status string `json:"status"`
}
