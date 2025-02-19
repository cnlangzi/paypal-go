package paypal

// PaymentSource represents different payment sources that can be used in an order.
// It is used to specify the type of payment method, such as PayPal, card, or token.
// This allows the order processing system to handle payments accordingly.
type PaymentSource string

const (
	// PaymentSourcePaypal Indicates that PayPal Wallet is the payment source.
	// Main use of this selection is to provide additional instructions associated with this choice like vaulting.
	PaymentSourcePaypal = PaymentSource("paypal")
	// PaymentSourceCard The payment card to use to fund a payment. Can be a credit or debit card.
	PaymentSourceCard = PaymentSource("card")
	// PaymentSourceToken The tokenized payment source to fund a payment.
	PaymentSourceToken = PaymentSource("token")
	// PaymentSourceBancontact Bancontact is the most popular online payment in Belgium. https://www.bancontact.com/
	PaymentSourceBancontact = PaymentSource("bancontact")
	// PaymentSourceBlik BLIK is a mobile payment system, created by Polish Payment Standard in order to allow millions of users
	// to pay in shops, payout cash in ATMs and make online purchases and payments.
	PaymentSourceBlik = PaymentSource("blik")
	// PaymentSourceEPS The eps transfer is an online payment method developed by many Austrian banks.
	PaymentSourceEPS = PaymentSource("eps")
	// PaymentSourceIDEAL The Dutch payment method iDEAL is an online payment method that enables consumers to pay online through their own bank.
	PaymentSourceIDEAL = PaymentSource("ideal")
	//PaymentSourceMyBank TMyBank is an e-authorisation solution which enables safe digital payments and identity authentication
	// through a consumer’s own online banking portal or mobile application.
	PaymentSourceMyBank = PaymentSource("mybank")
	// PaymentSourceP24 P24 (Przelewy24) is a secure and fast online bank transfer service linked to all the major banks in Poland.
	PaymentSourceP24 = PaymentSource("p24")
	// PaymentSourceSofort SOFORT Banking is a real-time bank transfer payment method that buyers use to transfer funds directly to merchants
	// from their bank accounts.
	PaymentSourceSofort = PaymentSource("sofort")
	//PaymentSourceTrustly 	Trustly is a payment method that allows customers to shop and pay from their bank account.
	PaymentSourceTrustly = PaymentSource("trustly")
	// PaymentSourceApplePay ApplePay payment source, allows buyer to pay using ApplePay, both on Web as well as on Native.
	PaymentSourceApplePay = PaymentSource("apple_pay")
	// PaymentSourceGooglePay Google Pay payment source, allows buyer to pay using Google Pay.
	PaymentSourceGooglePay = PaymentSource("google_pay")
	// PaymentSourceVenmo Information needed to indicate that Venmo is being used to fund the payment.
	PaymentSourceVenmo = PaymentSource("venmo")
)
