package yuansfer

import (
	"encoding/json"
)

type Securepay struct {
	MerchantNo   string `json:"merchantNo"`
	StoreNo      string `json:"storeNo"`
	VerifySign   string `json:"verifySign"`
	Currency     string `json:"currency"`
	Amount       string `json:"amount"`
	Vendor       string `json:"vendor"`
	Reference    string `json:"reference"`
	IpnUrl       string `json:"ipnUrl"`
	CallbackUrl  string `json:"callbackUrl"`
	Description  string `json:"description"`
	Note         string `json:"note"`
	Terminal     string `json:"terminal"`
	Timeout      string `json:"timeout"`
	Version      string `json:"version"`
	RmbAmount    string `json:"rmbAmount"`
	StoreAdminNo string `json:"storeAdminNo"`
	GoodsInfo    string `json:"goodsInfo"`
}

type GoodsInfomation struct {
	GoodsName string `json:"goods_name"`
	Quantity  string `json:"quantity"`
}

func (s Securepay) PostToYuansfer(token string) (string, error) {
	values := generateValues(s, token)
	securepayeUrl := yuansferHost + YuansferApi.OnlinePayment
	return postToYuansfer(securepayeUrl, values)
}

func (s *Securepay) Format(goodsInfos []GoodsInfomation) (err error) {
	if nil == goodsInfos {
		return
	}

	if bytes, err := json.Marshal(goodsInfos); nil == err {
		s.GoodsInfo = string(bytes)
	}

	return
}
