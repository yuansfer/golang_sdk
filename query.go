package yuansfer

// To get transaction details
// by ID of the transaction in the merchant’s system.
type Query struct {
	MerchantNo  string `json:"merchantNo"`
	StoreNo     string `json:"storeNo"`
	VerifySign  string `json:"verifySign"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	Vendor      string `json:"vendor"`
	Reference   string `json:"reference"`
	IpnURL      string `json:"ipnUrl"`
	CallbackURL string `json:"callbackUrl"`
	Description string `json:"description"`
	Note        string `json:"note"`
	Terminal    string `json:"terminal"`
	Timeout     string `json:"timeout"`
	Version     string `json:"version"`
}

//Send request to Yuansfer service
func (q Query) PostToYuansfer() (string, error) {
	values := generateValues(q, YuansferApi.Token.SecurepayToken)
	queryURL := yuansferHost + YuansferApi.OnlineQuery
	return postToYuansfer(queryURL, values)
}

//Response from the Yuansfer service.
type QryResponse struct {
	Result Ret    `json:"result"`
	RetMsg string `json:"ret_msg"`
}

//Response from the Yuansfer service.
type Ret struct {
	Reference  string       `json:"reference"`
	YuansferId string       `json:"yuansferId"`
	Amount     string       `json:"amount"`
	Status     string       `json:"status"`
	RefundInfo []refundInfo `json:"refundInfo"`
	Currency   string       `json:"currency"`
}

//Response from the Yuansfer service.
type refundInfo struct {
	RefundYuansferId string `json:"refundYuansferId"`
	RefundAmount     string `json:"refundAmount"`
}
