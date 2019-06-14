package yuansfer

//Reverse to cancel a transaction
// when the PAY API call did NOT return a clear result.
// - If the payment result of that trasaction is failure,
// the Yuansfer system will cancel the transaction.
// - If the payment result of that transaction is success,
// the Yuansfer system will refund the amount of the transaction.
type Reverse struct {
	MerchantNo    string `json:"merchantNo"`
	StoreNo       string `json:"storeNo"`
	TransactionNo string `json:"transactionNo"`
	VoidAmount    string `json:"voidAmount"`
	Reference     string `json:"reference"`
}

//PostToYuansfer is uesed to send request to the Yuansfer service
func (r Reverse) PostToYuansfer() (string, error) {
	values := generateValues(r, YuansferAPI.Token.InstoreToken)
	reverseURL := yuansferHost + YuansferAPI.InstoreReverse
	return postToYuansfer(reverseURL, values)
}

//QryReverse is the response from the Yuansfer service.
type QryReverse struct {
	RetCode string     `json:"ret_code"`
	RetMsg  string     `json:"ret_msg"`
	Result  ReverseRet `json:"refundTransaction"`
}

//ReverseRet is the response from the Yuansfer service.
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
	Reference             string `json:"reference"`
}
