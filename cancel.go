package yuansfer

//Cancel to cancel/reverse a transaction
type Cancel struct {
	GroupNo       string `json:"merGroupNo,omitempty"`
	MerchantNo    string `json:"merchantNo"`
	StoreNo       string `json:"storeNo"`
	VerifySign    string `json:"verifySign"`
	Reference     string `json:"reference,omitempty"`
	TransactionNo string `json:"transactionNo,omitempty"`
}

//PostToYuansfer is for api call to the Yuansfer service
func (c Cancel) PostToYuansfer() ([]byte, error) {
	(&c).LoadCredentials()
	c.VerifySign = getSignature(c)
	return postToYuansfer(URICancel, c)
}

func (c *Cancel) LoadCredentials() {
	c.GroupNo = cfg.Credential.PartnerNo
	c.MerchantNo = cfg.Credential.MerchantNo
	c.StoreNo = cfg.Credential.StoreNo
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
