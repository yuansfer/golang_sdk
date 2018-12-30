package controllers

import (
	"fmt"
	"html/template"
	"time"

	yuan "github.com/yuansfer/golang_sdk"
)

const (
	//securepay token
	yuansferToken = "3b581909d6294e988887838865e0bb76"
	//instore token
	instoreToken = "59600f2a9ad644c6a9570233560cc94e"
)

type HomeController struct {
	controller
}

func (this *HomeController) Get() {
	this.Data["IsPay"] = true
	this.TplName = "submit.tpl"
}

func (this *HomeController) Post() {
	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	amt := this.Input().Get("amt")
	vendor := this.Input().Get("vendor")
	rmbAmt := this.Input().Get("rmbAmt")
	reference := this.Input().Get("reference")
	description := this.Input().Get("description")
	note := this.Input().Get("note")
	terminal := this.Input().Get("terminal")
	ipnUrl := this.Input().Get("ipnUrl")
	callbackUrl := this.Input().Get("callbackUrl")
	token := this.Input().Get("token")

	if "" == terminal {
		terminal = "ONLINE"
	}
	if "" == reference {
		reference = fmt.Sprintf("seq_%d", time.Now().Unix())
	}
	if "" == token {
		token = yuansferToken
	}
	req := &yuan.Securepay{
		MerchantNo:  merchantNo,
		StoreNo:     storeNo,
		Currency:    "USD",
		Amount:      amt,
		RmbAmount:   rmbAmt,
		Vendor:      vendor,
		Reference:   reference,
		IpnUrl:      ipnUrl,
		CallbackUrl: callbackUrl,
		Description: description,
		Note:        note,
		Terminal:    terminal,
		Timeout:     "120",
	}

	goods := this.Input().Get("goods")
	quantity := this.Input().Get("quantity")
	if quantity != "" && goods != "" {
		goodsInfos := []yuan.GoodsInfomation{
			yuan.GoodsInfomation{
				GoodsName: goods,
				Quantity:  quantity,
			},
		}
		_ = req.Format(goodsInfos)
	}

	resp, err := req.PostToYuansfer(token)
	if err != nil {
		fmt.Println(err)
		this.Ctx.WriteString("something wrong happened")
	}
	t := template.New("secury pay template")
	t, _ = t.Parse(resp)
	_ = t.Execute(this.Ctx.ResponseWriter, resp)

	return
}
