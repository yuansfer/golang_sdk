package yuansfer

type Process struct {
	GroupNo    string `json:"merGroupNo"`
	MerchantNo string `json:"merchantNo"`
	StoreNo    string `json:"storeNo"`
	VerifySign string `json:"verifySign"`
	// PaymentMethodNonce received from user end.
	PaymentMethodNonce string `json:"paymentMethodNonce"`
	// PaymentMethod indicates which payment method client selected,
	// accessible values: paypal_account, venmo_account, credit_card, android_pay_card, apple_pay_card
	PaymentMethod string `json:"paymentMethod"`
	TransactionNo string `json:"transactionNo"`
	CustomerNo    string `json:"customerNo,omitempty"`
	// DeviceData received from user end provided by Braintree JS-SDK
	DeviceData string `json:"deviceData"`
}

//PostToYuansfer is for api call to the Yuansfer gateway
func (p Process) PostToYuansfer() ([]byte, error) {
	(&p).LoadCredentials()
	p.VerifySign = getSignature(p)
	return postToYuansfer(URIYipApp, p)
}

func (p *Process) LoadCredentials() {
	p.GroupNo = cfg.Credential.PartnerNo
	p.MerchantNo = cfg.Credential.MerchantNo
	p.StoreNo = cfg.Credential.StoreNo
}
