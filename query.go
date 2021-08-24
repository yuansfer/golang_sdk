package yuansfer

// Query obtain transaction details
// by ID of the transaction in the merchant’s system.
type Query struct {
	GroupNo       string `json:"merGroupNo,omitempty"`
	MerchantNo    string `json:"merchantNo"`
	StoreNo       string `json:"storeNo"`
	VerifySign    string `json:"verifySign"`
	Reference     string `json:"reference,omitempty"`
	TransactionNo string `json:"transactionNo,omitempty"`
}

//PostToYuansfer is for api call to the Yuansfer gateway
func (q Query) PostToYuansfer() ([]byte, error) {
	(&q).LoadCredentials()
	q.VerifySign = getSignature(q)
	return postToYuansfer(URIRetrieve, q)
}

func (q *Query) LoadCredentials() {
	q.GroupNo = cfg.Credential.PartnerNo
	q.MerchantNo = cfg.Credential.MerchantNo
	q.StoreNo = cfg.Credential.StoreNo
}

//QryResponse is the response from the Yuansfer service.
type QryResponse struct {
	Result Ret    `json:"result"`
	RetMsg string `json:"ret_msg"`
}

//Ret is the response from the Yuansfer service.
type Ret struct {
	Reference  string       `json:"reference"`
	YuansferID string       `json:"yuansferId"`
	Amount     string       `json:"amount"`
	Status     string       `json:"status"`
	RefundInfo []refundInfo `json:"refundInfo"`
	Currency   string       `json:"currency"`
}

type refundInfo struct {
	RefundYuansferID string `json:"refundYuansferId"`
	RefundAmount     string `json:"refundAmount"`
}
