package yuansfer

type Refund struct {
	MerchantNo       string `json:"merchantNo"`
	StoreNo          string `json:"storeNo"`
	ManagerAccountNo string `json:"managerAccountNo"`
	Password         string `json:"password"`
	Reference        string `json:"reference"`
	Amount           string `json:"amount"`
	RmbAmount        string `json:"rmbAmount"`
	Currency         string `json:"currency"`
	VerifySign       string `json:"verifySign"`
	Version          string `json:"version"`
}

func (r Refund) PostToYuansfer() (string, error) {

	r.Password = md5Token(YuansferApi.PwdPre + r.Password)

	values := generateValues(r, YuansferApi.Token.SecurepayToken)
	refundUrl := yuansferHost + YuansferApi.OnlineRefund
	return postToYuansfer(refundUrl, values)
}

type RefundResponse struct {
	Result  RefundBody `json:"result"`
	RetMsg  string     `json:"ret_msg"`
	RetCode string     `json:"ret_code"`
}

type RefundBody struct {
	OldTransactionId    string `json:"oldTransactionId"`
	RefundTransactionId string `json:"refundTransactionId"`
	Reference           string `json:"reference"`
	Status              string `json:"status"`
}
