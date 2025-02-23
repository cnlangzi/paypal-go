package paypal

type Status string

const (
	StatusCreated   Status = "CREATED"
	StatusCompleted Status = "COMPLETED"

	StatusEligible           Status = "ELIGIBLE"
	StatusInstrumentDeclined Status = "INSTRUMENT_DECLINED"

	StatusFailed Status = "FAILED"

	StatusRecoverable Status = "RECOVERABLE"

	// OrderStatusApproved   Status = "APPROVED"
	// OrderStatusCaptured   Status = "CAPTURED"
	// OrderStatusExpired    Status = "EXPIRED"
	// OrderStatusFailed     Status = "FAILED"
	// OrderStatusInReview   Status = "IN_REVIEW"
	// OrderStatusPartially  Status = "PARTIALLY_PAID"
	// OrderStatusProcessing Status = "PROCESSING"
	// OrderStatusRefunded   Status = "REFUNDED"
	// OrderStatusRevoked    Status = "REVOKED"
	// OrderStatusVoided     Status = "VOIDED"
)

type PaypalError struct {
	Name       string         `json:"name"`
	Message    string         `json:"message"`
	StatusCode int            `json:"-"`
	DebugID    string         `json:"debug_id"`
	Details    []PaypalDetail `json:"details"`
	Links      []Link         `json:"links"`
}

func (e *PaypalError) Error() string {
	return e.Message
}

type PaypalDetail struct {
	Field       string `json:"field"`
	Value       string `json:"value"`
	Location    string `json:"location"`
	Issue       string `json:"issue"`
	Description string `json:"description"`
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}
