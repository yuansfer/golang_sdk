package yuansfer

import (
	"encoding/json"
)

//Use the securepay() API to pay an order.
type Securepay struct {
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
}

//JSON encoded string of an array of items that the customer purchases
//from the merchant. Special characters are not supported.
// e.g.: [{"goods_name":"name1","quantity":"quantity1"},{"goods_name":"name2","quantity":"quantity2"}]
type GoodsInfomation struct {
	GoodsName string `json:"goods_name"`
	Quantity  string `json:"quantity"`
}

//Send request to Yuansfer service
func (s Securepay) PostToYuansfer() (string, error) {
	values := generateValues(s, YuansferApi.Token.SecurepayToken)
	securepayeURL := yuansferHost + YuansferApi.OnlinePayment
	return postToYuansfer(securepayeURL, values)
}

//Change GoodsInfo to string.
func (s *Securepay) Format(goodsInfos []GoodsInfomation) (err error) {
	if nil == goodsInfos {
		return
	}

	if bytes, err := json.Marshal(goodsInfos); nil == err {
		s.GoodsInfo = string(bytes)
	}

	return
}
