package yuansfer

type Query struct {
	MerchantNo  string `json:"merchantNo"`
	StoreNo     string `json:"storeNo"`
	VerifySign  string `json:"verifySign"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	Vendor      string `json:"vendor"`
	Reference   string `json:"reference"`
	IpnUrl      string `json:"ipnUrl"`
	CallbackUrl string `json:"callbackUrl"`
	Description string `json:"description"`
	Note        string `json:"note"`
	Terminal    string `json:"terminal"`
	Timeout     string `json:"timeout"`
	Version     string `json:"version"`
}

func (q Query) PostToYuansfer() (string, error) {
	values := generateValues(q, YuansferApi.Token.SecurepayToken)
	queryUrl := yuansferHost + YuansferApi.OnlineQuery
	return postToYuansfer(queryUrl, values)
}

type QryResponse struct {
	Result Ret    `json:"result"`
	RetMsg string `json:"ret_msg"`
}

type Ret struct {
	Reference  string       `json:"reference"`
	YuansferId string       `json:"yuansferId"`
	Amount     string       `json:"amount"`
	Status     string       `json:"status"`
	RefundInfo []refundInfo `json:"refundInfo"`
	Currency   string       `json:"currency"`
}

type refundInfo struct {
	RefundYuansferId string `json:"refundYuansferId"`
	RefundAmount     string `json:"refundAmount"`
}