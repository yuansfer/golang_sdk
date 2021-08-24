package yuansfer

//CreateQrcode is uesed to create transaction and get QR code
//for customers to scan to pay
type CreateQrcode struct {
	GroupNo      string `json:"merGroupNo,omitempty"`
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

//PostToYuansfer is for api call to the Yuansfer gateway
func (s CreateQrcode) PostToYuansfer() ([]byte, error) {
	(&s).LoadCredentials()
	s.VerifySign = getSignature(s)
	return postToYuansfer(URIPosCreateQRC, s)
}

func (s *CreateQrcode) LoadCredentials() {
	s.GroupNo = cfg.Credential.PartnerNo
	s.MerchantNo = cfg.Credential.MerchantNo
	s.StoreNo = cfg.Credential.StoreNo
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
