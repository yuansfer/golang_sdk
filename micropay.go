package yuansfer

type Micropay struct {
	MerchantNo   string `json:"merchantNo"`
	StoreNo      string `json:"storeNo"`
	VerifySign   string `json:"verifySign"`
	Currency     string `json:"currency"`
	Amount       string `json:"amount"`
	Vendor       string `json:"vendor"`
	Reference    string `json:"reference"`
	IpnUrl       string `json:"ipnUrl"`
	Description  string `json:"description"`
	Note         string `json:"note"`
	Terminal     string `json:"terminal"`
	Timeout      string `json:"timeout"`
	Version      string `json:"version"`
	RmbAmount    string `json:"rmbAmount"`
	StoreAdminNo string `json:"storeAdminNo"`
	Openid       string `json:"openid"`
}

func (s Micropay) PostToYuansfer(token string) (string, error) {
	values := generateValues(s, token)
	requestUrl := yuansferHost + YuansferApi.Micropay
	return postToYuansfer(requestUrl, values)
}
