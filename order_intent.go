package paypal

// OrderIntent represents the intent of the order. It can be either "CAPTURE"
// or "AUTHORIZE". If it is "CAPTURE", the order is automatically captured and
// the payment is taken. If it is "AUTHORIZE", the order is only authorized and
// the payment is not taken until the order is captured.
//
// https://developer.paypal.com/docs/api/orders/v2/#orders_create!ct=application/json&path=intent&t=request
type OrderIntent string

const (
	// OrderIntentCapture The merchant intends to capture payment immediately after the customer makes a payment.
	OrderIntentCapture = OrderIntent("CAPTURE")
	// OrderIntentAuthorize The merchant intends to authorize a payment and place funds on hold after the customer makes a payment.
	// Authorized payments are best captured within three days of authorization but are available to capture for up to 29 days.
	// After the three-day honor period, the original authorized payment expires and you must re-authorize the payment.
	// You must make a separate request to capture payments on demand.
	// This intent is not supported when you have more than one purchase_unit within your order.
	OrderIntentAuthorize = OrderIntent("AUTHORIZE")
)
