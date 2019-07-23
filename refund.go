package yuansfer

// Refund To refund payments.
type Refund struct {
	MerchantNo       string `json:"merchantNo"`
	StoreNo          string `json:"storeNo"`
	ManagerAccountNo string `json:"managerAccountNo"`
	Password         string `json:"password"`
	Reference        string `json:"reference"`
	Amount           string `json:"amount"`
	RmbAmount        string `json:"rmbAmount"`
	Currency         string `json:"currency"`
	VerifySign       string `json:"verifySign"`
	Version          string `json:"version"`
	TransactionNo    string `json:"transactionNo"`
}

//PostToYuansfer is uesed to send request to the Yuansfer service
func (r Refund) PostToYuansfer() (string, error) {
	r.Password = md5Token(YuansferAPI.PwdPre + r.Password)
	values := generateValues(r, YuansferAPI.Token.SecurepayToken)
	refundURL := yuansferHost + YuansferAPI.OnlineRefund
	return postToYuansfer(refundURL, values)
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
