package yuansfer

//InStoreAdd initiates a Barcode/QR Code Payment request and creates a transaction order.
type InStoreAdd struct {
	GroupNo      string `json:"merGroupNo,omitempty"`
	MerchantNo   string `json:"merchantNo"`
	StoreNo      string `json:"storeNo"`
	VerifySign   string `json:"verifySign"`
	Amount       string `json:"amount"`
	StoreAdminNo string `json:"storeAdminNo"`
	IpnUrl       string `json:"ipnUrl"`
	Reference    string `json:"reference"`
}

//PostToYuansfer is for api call to the Yuansfer gateway
func (s InStoreAdd) PostToYuansfer() ([]byte, error) {
	(&s).LoadCredentials()
	s.VerifySign = getSignature(s)
	return postToYuansfer(URIPosAdd, s)
}

func (s *InStoreAdd) LoadCredentials() {
	s.GroupNo = cfg.Credential.PartnerNo
	s.MerchantNo = cfg.Credential.MerchantNo
	s.StoreNo = cfg.Credential.StoreNo
}

//AddResponse is the Response from the Yuansfer service.
type AddResponse struct {
	Result      string `json:"ret_code"`
	RetMsg      string `json:"ret_msg"`
	Transaction AddRet `json:"transaction"`
}

//AddRet is the Response from the Yuansfer service.
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
