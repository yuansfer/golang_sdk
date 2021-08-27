package yuansfer

import (
	"encoding/json"
)

//SecurePay contains the attributes of online payment request.
type SecurePay struct {
	GroupNo    string `json:"merGroupNo,omitempty"`
	MerchantNo string `json:"merchantNo"`
	StoreNo    string `json:"storeNo"`
	VerifySign string `json:"verifySign"`
	Currency   string `json:"currency"`
	// SettleCurrency refers to settlement currency, default value is USD
	SettleCurrency string `json:"settleCurrency,omitempty"`
	Amount         string `json:"amount"`
	Vendor         string `json:"vendor"`
	Reference      string `json:"reference"`
	Terminal       string `json:"terminal"`
	IpnURL         string `json:"ipnUrl"`
	CallbackURL    string `json:"callbackUrl,omitempty"`
	Note           string `json:"note,omitempty"`
	Description    string `json:"description,omitempty"`
	Timeout        string `json:"timeout,omitempty"`
	GoodsInfo      string `json:"goodsInfo,omitempty"`
	CreditType     string `json:"creditType,omitempty"`
	CustomerNo     string `json:"customerNo,omitempty"`
	OSType         string `json:"osType,omitempty"` // IOS, ANDROID available
}

//AddResponse is the Response from the Yuansfer service.
type SecurePayResponse struct {
	Result   string       `json:"ret_code"`
	RetMsg   string       `json:"ret_msg"`
	Response SecurePayRet `json:"result"`
}

//AddRet is the Response from the Yuansfer service.
type SecurePayRet struct {
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	TransactionNo  string `json:"transactionNo"`
	Reference      string `json:"reference"`
	CashierURL     string `json:"cashierUrl"`
	SettleCurrency string `json:"settleCurrency"`
}

//GoodsInformation is a JSON encoded string of an array of items that the customer purchases
//from the merchant. Special characters are not supported.
// e.g.: [{"goods_name":"name1","quantity":"quantity1"},{"goods_name":"name2","quantity":"quantity2"}]
type GoodsInformation struct {
	GoodsName string `json:"goods_name"`
	Quantity  string `json:"quantity"`
}

//PostToYuansfer is for api call to the Yuansfer gateway
func (s SecurePay) PostToYuansfer() ([]byte, error) {
	(&s).LoadCredentials()
	s.VerifySign = getSignature(s)
	return postToYuansfer(URISecurePay, s)
}

//Format changes GoodsInfo to string.
func (s *SecurePay) Format(goodsInfos []GoodsInformation) (err error) {
	if nil == goodsInfos {
		return
	}

	if bytes, err := json.Marshal(goodsInfos); nil == err {
		s.GoodsInfo = string(bytes)
	}
	return
}

func (s *SecurePay) LoadCredentials() {
	s.GroupNo = cfg.Credential.PartnerNo
	s.MerchantNo = cfg.Credential.MerchantNo
	s.StoreNo = cfg.Credential.StoreNo
}
