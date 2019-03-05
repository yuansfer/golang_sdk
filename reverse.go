package yuansfer

type Reverse struct {
	MerchantNo    string `json:"merchantNo"`
	StoreNo       string `json:"storeNo"`
	TransactionNo string `json:"transactionNo"`
}

func (r Reverse) PostToYuansfer() (string, error) {
	values := generateValues(r, YuansferApi.Token.InstoreToken)
	reverseUrl := yuansferHost + YuansferApi.InstoreReverse
	return postToYuansfer(reverseUrl, values)
}

type QryReverse struct {
	RetCode string     `json:"ret_code"`
	RetMsg  string     `json:"ret_msg"`
	Result  ReverseRet `json:"refundTransaction"`
}

type ReverseRet struct {
	ExchangeRate          string `json:"exchangeRate"`
	Currency              string `json:"currency"`
	Amount                string `json:"amount"`
	CreateTime            string `json:"createTime"`
	MerchantNo            string `json:"merchantNo"`
	OriginalTransactionNo string `json:"originalTransactionNo"`
	PaymentTime           string `json:"paymentTime"`
	RefundAmount          string `json:"refundAmount"`
	StoreAdminNo          string `json:"storeAdminNo"`
	StoreNo               string `json:"storeNo"`
	TransactionNo         string `json:"transactionNo"`
	TransactionStatus     string `json:"transactionStatus"`
	TransactionType       string `json:"transactionType"`
}