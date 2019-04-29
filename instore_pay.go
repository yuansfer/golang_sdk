package yuansfer

//InstorePay for placeing an order in the Barcode/QR Code Payment.
type InstorePay struct {
	MerchantNo     string `json:"merchantNo"`
	StoreNo        string `json:"storeNo"`
	VerifySign     string `json:"verifySign"`
	PaymentBarcode string `json:"paymentBarcode"`
	StoreAdminNo   string `json:"storeAdminNo"`
	TransactionNo  string `json:"transactionNo"`
	Vendor         string `json:"vendor"`
}

//PostToYuansfer is uesed to send request to the Yuansfer service
func (s InstorePay) PostToYuansfer() (string, error) {
	values := generateValues(s, YuansferAPI.Token.InstoreToken)
	requestURL := yuansferHost + YuansferAPI.InstorePay
	return postToYuansfer(requestURL, values)
}

//PayResponse is the response from the Yuansfer service.
type PayResponse struct {
	Result      AddRet `json:"ret_code"`
	RetMsg      string `json:"ret_msg"`
	Transaction AddRet `json:"transaction"`
}

//PayRet is the response from the Yuansfer service.
type PayRet struct {
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
