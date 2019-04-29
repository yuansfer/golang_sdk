package yuansfer

//To initiate a Barcode/QR Code Payment request
//and create transaction order.
type InstoreAdd struct {
	MerchantNo   string `json:"merchantNo"`
	StoreNo      string `json:"storeNo"`
	VerifySign   string `json:"verifySign"`
	Amount       string `json:"amount"`
	StoreAdminNo string `json:"storeAdminNo"`
}

//Send request to Yuansfer service
func (s InstoreAdd) PostToYuansfer() (string, error) {
	values := generateValues(s, YuansferApi.Token.InstoreToken)
	requestURL := yuansferHost + YuansferApi.InstoreAdd
	return postToYuansfer(requestURL, values)
}

//Response from the Yuansfer service.
type AddResponse struct {
	Result      AddRet `json:"ret_code"`
	RetMsg      string `json:"ret_msg"`
	Transaction AddRet `json:"transaction"`
}

//Response from the Yuansfer service.
type AddRet struct {
	TransactionNo         string `json:"transactionNo"`
	OriginalTransactionNo string `json:"originalTransactionNo"`
	MerchantNo            string `json:"merchantNo"`
	StoreNo               string `json:"storeNo"`
	StoreAdminNo          string `json:"storeAdminNo"`
	Amount                string `json:"amount"`
	RefundAmount          string `json:"refundAmount"`
	TransactionType       string `json:"transactionType"`
	TransactionStatus     string `json:"transactionStatus"`
	Currency              string `json:"currency"`
	CreateTime            string `json:"createTime"`
	PaymentTime           string `json:"paymentTime"`
	ExchangeRate          string `json:"exchangeRate"`
	Vendor                string `json:"vendor"`
}
