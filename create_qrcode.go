package yuansfer

//CreateQrcode is uesed to create transaction and get QR code
//for customers to scan to pay
type CreateQrcode struct {
	MerchantNo   string `json:"merchantNo"`
	StoreNo      string `json:"storeNo"`
	VerifySign   string `json:"verifySign"`
	Amount       string `json:"amount"`
	StoreAdminNo string `json:"storeAdminNo"`
	Currency     string `json:"currency"`
	Vendor       string `json:"vendor"`
	Reference    string `json:"reference"`
	IpnURL       string `json:"ipnUrl"`
}

//PostToYuansfer is uesed to send request to the Yuansfer service
func (s CreateQrcode) PostToYuansfer() (string, error) {
	values := generateValues(s, YuansferAPI.Token.InstoreToken)
	requestURL := yuansferHost + YuansferAPI.InstoreCreateQrcode
	return postToYuansfer(requestURL, values)
}

//QrcodeRet is the Response from the Yuansfer service
type QrcodeRet struct {
	Reference     string `json:"reference"`
	RectCode      string `json:"ret_code"`
	DeepLink      string `json:"deepLink"`
	TransactionNo string `json:"transactionNo"`
	TimeOut       string `json:"timeout"`
	QrcodeURL     string `json:"qrcodeUrl"`
}
