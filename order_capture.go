package paypal

type Shipping struct {
	Name    Name    `json:"name,omitzero"`
	Address Address `json:"address,omitzero"`
}

type CapturedOrder struct {
	ID     string `json:"id"`
	Status Status `json:"status"`
	Links  []Link `json:"links"`

	PaymentSource PaymentSource `json:"payment_source,omitzero"`

	PurchaseUnits []CapturedOrderPurchaseUnit `json:"purchase_units,omitzero"`

	Payer Payer `json:"payer,omitzero"`
}

type CapturedOrderPurchaseUnit struct {
	ReferenceID string   `json:"reference_id,omitzero"`
	Shipping    Shipping `json:"shipping,omitzero"`

	Payments CapturedOrderPayments `json:"payments,omitzero"`
}

type CapturedOrderPayments struct {
	Captures []CapturedOrderPayment `json:"captures"`
}

type CapturedOrderPayment struct {
	ID        string     `json:"id"`
	Status    Status     `json:"status"`
	Amount    AmountItem `json:"amount"`
	InvoiceID string     `json:"invoice_id,omitzero"`

	FinalCapture              bool                      `json:"final_capture,omitzero"`
	SellerProtection          SellerProtection          `json:"seller_protection,omitzero"`
	SellerReceivableBreakdown SellerReceivableBreakdown `json:"seller_receivable_breakdown,omitzero"`
	Links                     []Link                    `json:"links,omitzero"`
	CreateTime                string                    `json:"create_time,omitzero"`
	UpdateTime                string                    `json:"update_time,omitzero"`
}

type SellerProtection struct {
	Status            Status   `json:"status"`
	DisputeCategories []string `json:"dispute_categories,omitzero"`
}

type SellerReceivableBreakdown struct {
	GrossAmount AmountItem `json:"gross_amount,omitzero"`
	NetAmount   AmountItem `json:"net_amount,omitzero"`
	PaypalFee   AmountItem `json:"paypal_fee,omitzero"`
}
