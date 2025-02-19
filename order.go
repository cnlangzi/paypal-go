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

// CreateOrderRequest represents the request body of the Create Order API.
// It contains the intent of the order, the purchase units, and the payment source.
type CreateOrderRequest struct {
	Intent        OrderIntent           `json:"intent"`
	PurchaseUnits []PurchaseUnit        `json:"purchase_units"`
	PaymentSource map[PaymentSource]any `json:"payment_source,omitempty"`
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

// WithPaymentSource sets the payment source for the CreateOrderRequest.
// It takes a PaymentSource and a value, and adds them to the PaymentSource map in the request.
func WithPaymentSource(name PaymentSource, value any) CreateOrderOption {
	return func(req *CreateOrderRequest) {
		req.PaymentSource[name] = value
	}
}

type Order struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Links  []Link `json:"links"`
}
