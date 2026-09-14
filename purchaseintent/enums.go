package purchaseintent

type Status string

const (
	StatusActive   Status = "active"
	StatusExpired  Status = "expired"
	StatusInactive Status = "inactive"
	StatusUsed     Status = "used"
)
