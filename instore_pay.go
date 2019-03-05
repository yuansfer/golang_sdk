package yuansfer

type InstorePay struct {
	MerchantNo     string `json:"merchantNo"`
	StoreNo        string `json:"storeNo"`
	VerifySign     string `json:"verifySign"`
	PaymentBarcode string `json:"paymentBarcode"`
	StoreAdminNo   string `json:"storeAdminNo"`
	TransactionNo  string `json:"transactionNo"`
	Vendor         string `json:"vendor"`
}

func (s InstorePay) PostToYuansfer() (string, error) {
	values := generateValues(s, YuansferApi.Token.InstoreToken)
	requestUrl := yuansferHost + YuansferApi.InstorePay
	return postToYuansfer(requestUrl, values)
}

type PayResponse struct {
	Result      AddRet `json:"ret_code"`
	RetMsg      string `json:"ret_msg"`
	Transaction AddRet `json:"transaction"`
}

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
