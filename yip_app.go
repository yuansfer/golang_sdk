package yuansfer

//Prepay does the mobile payment for the In-APP and/or WeChat Mini program solutions.
type Prepay struct {
	GroupNo        string `json:"merGroupNo,omitempty"`
	MerchantNo     string `json:"merchantNo"`
	StoreNo        string `json:"storeNo"`
	VerifySign     string `json:"verifySign"`
	Currency       string `json:"currency"`
	SettleCurrency string `json:"settleCurrency"`
	Amount         string `json:"amount"`
	Vendor         string `json:"vendor"` // Vendor, possible values are alipay, wechatpay
	Reference      string `json:"reference"`
	IpnURL         string `json:"ipnUrl"` // InpURL The instant Payment Notification Handler URL. IPN URL must be secure.
	Description    string `json:"description"`
	Note           string `json:"note"`
	Terminal       string `json:"terminal"` // Terminal including MINIPROGRAM, APP
	Timeout        string `json:"timeout"`
	Openid         string `json:"openid,omitempty"`
}

//PostToYuansfer is for api call to the Yuansfer gateway
func (p Prepay) PostToYuansfer() ([]byte, error) {
	(&p).LoadCredentials()
	p.VerifySign = getSignature(p)
	return postToYuansfer(URIYipApp, p)
}

func (p *Prepay) LoadCredentials() {
	p.GroupNo = cfg.Credential.PartnerNo
	p.MerchantNo = cfg.Credential.MerchantNo
	p.StoreNo = cfg.Credential.StoreNo
}
