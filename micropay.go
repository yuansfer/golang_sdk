package yuansfer

//Micropay is for WeChat Mini-Program Payment
type Micropay struct {
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

//PostToYuansfer is uesed to send request to the Yuansfer service
func (s Micropay) PostToYuansfer() (string, error) {
	values := generateValues(s, YuansferAPI.Token.MicropayToken)
	requestURL := yuansferHost + YuansferAPI.Micropay
	return postToYuansfer(requestURL, values)
}
