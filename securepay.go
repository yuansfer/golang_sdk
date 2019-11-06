package yuansfer

import (
	"encoding/json"
)

//Securepay contains the parameters of online payment request.
type Securepay struct {
	GroupNo      string `json:"merGroupNo"`
	MerchantNo   string `json:"merchantNo"`
	StoreNo      string `json:"storeNo"`
	VerifySign   string `json:"verifySign"`
	Currency     string `json:"currency"`
	Amount       string `json:"amount"`
	Vendor       string `json:"vendor"`
	Reference    string `json:"reference"`
	IpnURL       string `json:"ipnUrl"`
	CallbackURL  string `json:"callbackUrl"`
	Description  string `json:"description"`
	Note         string `json:"note"`
	Terminal     string `json:"terminal"`
	Timeout      string `json:"timeout"`
	Version      string `json:"version"`
	RmbAmount    string `json:"rmbAmount"`
	StoreAdminNo string `json:"storeAdminNo"`
	GoodsInfo    string `json:"goodsInfo"`
	CreditType   string `json:"creditType"`
	PaymentCount string `json:"paymentCount"`
	Frequency    string `json:"frequency"`
	CustomerNo   string `json:"customerNo"`
}

//GoodsInfomation is a JSON encoded string of an array of items that the customer purchases
//from the merchant. Special characters are not supported.
// e.g.: [{"goods_name":"name1","quantity":"quantity1"},{"goods_name":"name2","quantity":"quantity2"}]
type GoodsInfomation struct {
	GoodsName string `json:"goods_name"`
	Quantity  string `json:"quantity"`
}

//PostToYuansfer is uesed to send request to the Yuansfer service
func (s Securepay) PostToYuansfer() (string, error) {
	values := generateValues(s, YuansferAPI.Token.SecurepayToken)
	securepayeURL := yuansferHost + YuansferAPI.OnlinePayment
	return postToYuansfer(securepayeURL, values)
}

//Format changes GoodsInfo to string.
func (s *Securepay) Format(goodsInfos []GoodsInfomation) (err error) {
	if nil == goodsInfos {
		return
	}

	if bytes, err := json.Marshal(goodsInfos); nil == err {
		s.GoodsInfo = string(bytes)
	}

	return
}
