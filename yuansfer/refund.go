package yuansfer

type Refund struct {
	MerchantNo       string `json:"merchantNo"`
	StoreNo          string `json:"storeNo"`
	ManagerAccountNo string `json:"managerAccountNo"`
	Password         string `json:"password"`
	Reference        string `json:"reference"`
	Amount           string `json:"amount"`
	Currency         string `json:"currency"`
	VerifySign       string `json:"verifySign"`
	Version          string `json:"version"`
}

func (r Refund) PostToYuansfer(token string) (string, error) {
	values := generateValues(r, token)
	refundUrl := yuansferHost + YuansferApi.OnlineRefund
	return postToYuansfer(refundUrl, values)
}
