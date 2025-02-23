package paypal

import "strconv"

func CreatePurchaseUnit(desc string, currencyCode string, amount float64, options ...PurchaseUnitOption) PurchaseUnit {
	u := PurchaseUnit{
		Description: desc,
		Amount: Amount{
			AmountItem: AmountItem{
				CurrencyCode: currencyCode,
				Amount:       amount,
				Value:        strconv.FormatFloat(amount, 'f', 2, 64),
			},
		},
	}

	for _, option := range options {
		option(&u)
	}

	return u
}

type PurchaseUnitOption func(*PurchaseUnit)

func WithPurchaseUnitCustomerID(customerID string) PurchaseUnitOption {
	return func(unit *PurchaseUnit) {
		unit.CustomerID = customerID
	}
}
func WithPurchaseUnitInvoiceID(invoiceID string) PurchaseUnitOption {
	return func(unit *PurchaseUnit) {
		unit.InvoiceID = invoiceID
	}
}

func WithPurchaseUnitShippingFee(fee float64) PurchaseUnitOption {
	return func(unit *PurchaseUnit) {

		unit.Amount.Breakdown = Breakdown{
			ItemAmount: AmountItem{
				CurrencyCode: unit.Amount.CurrencyCode,
				Amount:       unit.Amount.Amount,
				Value:        strconv.FormatFloat(unit.Amount.AmountItem.Amount, 'f', 2, 64),
			},
			Shipping: AmountItem{
				CurrencyCode: unit.Amount.CurrencyCode,
				Amount:       fee,
				Value:        strconv.FormatFloat(fee, 'f', 2, 64),
			},
		}

		unit.Amount.AmountItem.Amount += fee
		unit.Amount.AmountItem.Value = strconv.FormatFloat(unit.Amount.AmountItem.Amount, 'f', 2, 64)
	}
}

// PurchaseUnit is a purchase unit
// https://developer.paypal.com/docs/api/orders/v2/#orders_create!ct=application/json&path=purchase_units&t=request
type PurchaseUnit struct {

	//description	The purchase description.
	Description string `json:"description,omitzero"`
	// CustomerID The API caller-provided external ID. Used to reconcile client transactions with PayPal transactions.
	// Appears in transaction and settlement reports but is not visible to the payer.
	CustomerID string `json:"customer_id,omitzero"`
	// InvoiceID The API caller-provided external invoice number for this order.
	// Appears in both the payer's transaction history and the emails that the payer receives.
	InvoiceID string `json:"invoice_id,omitzero"`

	Amount Amount `json:"amount"`
	PurchaseUnitDetails
}

type PurchaseUnitDetails struct {
	// ReferenceID The API caller-provided external ID for the purchase unit.
	// Required for multiple purchase units when you must update the order through PATCH. If you omit
	// this value and the order contains only one purchase unit, PayPal sets this value to `default`.
	ReferenceID string `json:"reference_id,omitzero"`

	// SoftDescriptor The soft descriptor is the dynamic text used to construct the statement descriptor
	// that appears on a payer's card statement.
	SoftDescriptor string `json:"soft_descriptor,omitzero"`
}

// AmountItem is the amount of a purchase unit
type AmountItem struct {
	CurrencyCode string  `json:"currency_code"`
	Value        string  `json:"value"`
	Amount       float64 `json:"-"`
}

// Breakdown is the breakdown of the amount
type Breakdown struct {
	ItemAmount AmountItem `json:"item_amount"`
	Shipping   AmountItem `json:"shipping_amount"`
}

// Amount is the total amount of a purchase unit
type Amount struct {
	AmountItem
	Breakdown Breakdown `json:"breakdown,omitzero"`
}
