package yuansfer

//Prepay does the mobile payment for the In-APP and/or WeChat Mini program solutions.
type Prepay struct {
	GroupNo      string `json:"merGroupNo,omitempty"`
	MerchantNo   string `json:"merchantNo"`
	StoreNo      string `json:"storeNo"`
	VerifySign   string `json:"verifySign"`
	Currency     string `json:"currency"`
	Amount       string `json:"amount"`
	Vendor       string `json:"vendor"`
	Reference    string `json:"reference"`
	IpnURL       string `json:"ipnUrl"`
	Description  string `json:"description"`
	Note         string `json:"note"`
	Terminal     string `json:"terminal"`
	Timeout      string `json:"timeout"`
	Version      string `json:"version"`
	RmbAmount    string `json:"rmbAmount"`
	StoreAdminNo string `json:"storeAdminNo"`
	Openid       string `json:"openid"`
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
