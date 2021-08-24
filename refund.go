package yuansfer

// Refund To refund payments.
type Refund struct {
	GroupNo    string `json:"merGroupNo,omitempty"`
	MerchantNo string `json:"merchantNo"`
	StoreNo    string `json:"storeNo"`
	VerifySign string `json:"verifySign"`
	Amount     string `json:"refundAmount"`
	Currency   string `json:"currency"`
	// Reference indicates the payment transaction id.
	Reference     string `json:"reference,omitempty"`
	TransactionNo string `json:"transactionNo,omitempty"`
	// RefundReference used for refund transaction
	RefundReference string `json:"refundReference,omitempty"`
	SettleCurrency  string `json:"settleCurrency,omitempty"`
}

//PostToYuansfer is for api call to the Yuansfer gateway
func (r Refund) PostToYuansfer() ([]byte, error) {
	(&r).LoadCredentials()
	r.VerifySign = getSignature(r)
	return postToYuansfer(URIRefund, r)
}

func (r *Refund) LoadCredentials() {
	r.GroupNo = cfg.Credential.PartnerNo
	r.MerchantNo = cfg.Credential.MerchantNo
	r.StoreNo = cfg.Credential.StoreNo
}

//RefundResponse is the response from the Yuansfer service.
type RefundResponse struct {
	Result  RefundBody `json:"result"`
	RetMsg  string     `json:"ret_msg"`
	RetCode string     `json:"ret_code"`
}

//RefundBody is the response from the Yuansfer service.
type RefundBody struct {
	OldTransactionID    string `json:"oldTransactionId"`
	RefundTransactionID string `json:"refundTransactionId"`
	Reference           string `json:"reference"`
	Status              string `json:"status"`
}
