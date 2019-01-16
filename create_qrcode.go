package yuansfer

type CreateQrcode struct {
	MerchantNo   string `json:"merchantNo"`
	StoreNo      string `json:"storeNo"`
	VerifySign   string `json:"verifySign"`
	Amount       string `json:"amount"`
	StoreAdminNo string `json:"storeAdminNo"`
	Currency     string `json:"currency"`
	Vendor       string `json:"vendor"`
	Reference    string `json:"reference"`
	IpnUrl       string `json:"ipnUrl"`
}

func (s CreateQrcode) PostToYuansfer(token string) (string, error) {
	values := generateValues(s, token)
	requestUrl := yuansferHost + YuansferApi.InstoreCreateQrcode
	return postToYuansfer(requestUrl, values)
}

type QrcodeRet struct {
	Reference     string `json:"reference"`
	RectCode      string `json:"ret_code"`
	DeepLink      string `json:"deepLink"`
	TransactionNo string `json:"transactionNo"`
	TimeOut       string `json:"timeout"`
	QrcodeUrl     string `json:"qrcodeUrl"`
}
