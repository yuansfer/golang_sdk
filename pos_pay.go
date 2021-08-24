package yuansfer

//InStorePay to complete the QRC payment added previously.
type InStorePay struct {
	GroupNo        string `json:"merGroupNo,omitempty"`
	MerchantNo     string `json:"merchantNo"`
	StoreNo        string `json:"storeNo"`
	VerifySign     string `json:"verifySign"`
	PaymentBarcode string `json:"paymentBarcode"`
	StoreAdminNo   string `json:"storeAdminNo"`
	TransactionNo  string `json:"transactionNo"`
	Vendor         string `json:"vendor"`
}

//PostToYuansfer is for api call to the Yuansfer gateway
func (s InStorePay) PostToYuansfer() ([]byte, error) {
	(&s).LoadCredentials()
	s.VerifySign = getSignature(s)
	return postToYuansfer(URIPosPay, s)
}

func (s *InStorePay) LoadCredentials() {
	s.GroupNo = cfg.Credential.PartnerNo
	s.MerchantNo = cfg.Credential.MerchantNo
	s.StoreNo = cfg.Credential.StoreNo
}

//PayResponse is the response from the Yuansfer service.
type PayResponse struct {
	Result      string `json:"ret_code"`
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
