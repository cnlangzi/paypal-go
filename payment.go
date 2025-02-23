package paypal

// https://developer.paypal.com/docs/api/payments/v2/

//	{
//		"id": "0VF52814937998046",
//		"status": "CREATED",
//		"amount": {
//		"value": "10.99",
//		"currency_code": "USD"
//		},
//		"invoice_id": "INVOICE-123",
//		"seller_protection": {
//		"status": "ELIGIBLE",
//		"dispute_categories": [
//		"ITEM_NOT_RECEIVED",
//		"UNAUTHORIZED_TRANSACTION"
//		]
//		},
//		"payee": {
//		"email_address": "merchant@example.com",
//		"merchant_id": "7KNGBPH2U58GQ"
//		},
//		"expiration_time": "2017-10-10T23:23:45Z",
//		"create_time": "2017-09-11T23:23:45Z",
//		"update_time": "2017-09-11T23:23:45Z",
//		"links": [
//		{
//		"rel": "self",
//		"method": "GET",
//		"href": "https://api-m.paypal.com/v2/payments/authorizations/0VF52814937998046"
//		},
//		{
//		"rel": "capture",
//		"method": "POST",
//		"href": "https://api-m.paypal.com/v2/payments/authorizations/0VF52814937998046/capture"
//		},
//		{
//		"rel": "void",
//		"method": "POST",
//		"href": "https://api-m.paypal.com/v2/payments/authorizations/0VF52814937998046/void"
//		},
//		{
//		"rel": "reauthorize",
//		"method": "POST",
//		"href": "https://api-m.paypal.com/v2/payments/authorizations/0VF52814937998046/reauthorize"
//		}
//		]
//		}
type Payment struct {
	ID        string     `json:"id"`
	Status    string     `json:"status"`
	Amount    AmountItem `json:"amount"`
	InvoiceID string     `json:"invoice_id"`

	Payee            Payee            `json:"payee"`
	SellerProtection SellerProtection `json:"seller_protection"`

	ExpirationTime string `json:"expiration_time"` //		"expiration_time": "2017-10-10T23:23:45Z",
	CreateTime     string `json:"create_time"`     //		"create_time": "2017-09-11T23:23:45Z",
	UpdateTime     string `json:"update_time"`     //		"update_time": "2017-09-11T23:23:45Z",

	Links []Link `json:"links"`
}

type Payee struct {
	EmailAddress string `json:"email_address"`
	MerchantID   string `json:"merchant_id"`
}
