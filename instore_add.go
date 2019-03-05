package yuansfer

type InstoreAdd struct {
	MerchantNo   string `json:"merchantNo"`
	StoreNo      string `json:"storeNo"`
	VerifySign   string `json:"verifySign"`
	Amount       string `json:"amount"`
	StoreAdminNo string `json:"storeAdminNo"`
}

func (s InstoreAdd) PostToYuansfer() (string, error) {
	values := generateValues(s, YuansferApi.Token.InstoreToken)
	requestUrl := yuansferHost + YuansferApi.InstoreAdd
	return postToYuansfer(requestUrl, values)
}

type AddResponse struct {
	Result      AddRet `json:"ret_code"`
	RetMsg      string `json:"ret_msg"`
	Transaction AddRet `json:"transaction"`
}

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
