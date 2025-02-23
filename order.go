package paypal

// https://developer.paypal.com/docs/api/orders/v2/

// NewCreateOrderRequest create a new CreateOrderRequest with default values
func NewCreateOrderRequest(unit PurchaseUnit, options ...CreateOrderOption) *CreateOrderRequest {
	req := &CreateOrderRequest{
		Intent: OrderIntentCapture,
		PurchaseUnits: []PurchaseUnit{
			unit,
		},
	}

	for _, option := range options {
		option(req)
	}

	return req

}

// CreatedOrder represents a PayPal order that has been created.
// It contains the ID, status, and links related to the order.
type CreatedOrder struct {
	ID     string `json:"id"`     // The unique identifier for the order.
	Status Status `json:"status"` // The current status of the order.
	Links  []Link `json:"links"`  // Related links for the order actions.
}

// CreateOrderRequest represents the request body of the Create Order API.
// It contains the intent of the order, the purchase units, and the payment source.
type CreateOrderRequest struct {
	Intent        OrderIntent    `json:"intent"`
	PurchaseUnits []PurchaseUnit `json:"purchase_units"`
	PaymentSource PaymentSource  `json:"payment_source,omitzero"`
}

// CreateOrderOption is a function that take a pointer to a CreateOrderRequest and modify it.
// They are used to configure a CreateOrderRequest when calling the CreateOrder function.
type CreateOrderOption func(*CreateOrderRequest)

// WithOrderIntent sets the intent of the order to the given intent.
//
// It is a required field, and it can be either "CAPTURE" or "AUTHORIZE".
func WithOrderIntent(intent OrderIntent) CreateOrderOption {
	return func(req *CreateOrderRequest) {
		req.Intent = intent
	}
}

// WithPurchaseUnits sets the purchase units for the CreateOrderRequest.
//
// It takes a variadic argument of PurchaseUnit, and appends them to the PurchaseUnits field in the request.
func WithPurchaseUnit(units ...PurchaseUnit) CreateOrderOption {
	return func(req *CreateOrderRequest) {
		req.PurchaseUnits = append(req.PurchaseUnits, units...)
	}
}

// WithPaypalPaymentSources sets the payment source of the CreateOrderRequest to the given PayPalPaymentSource.
//
// It takes a PaypalPaymentSource as an argument, and sets the PaymentSource field in the request to it.
func WithPaypalPaymentSources(s PaypalPaymentSource) CreateOrderOption {
	return func(req *CreateOrderRequest) {
		req.PaymentSource.Paypal = s
	}
}
