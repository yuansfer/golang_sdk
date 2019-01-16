package controllers

import (
	"fmt"
	"html/template"
	"log"
	"time"

	yuan "github.com/yuansfer/golang_sdk"
)

const (
	miropayToken = "4dc2f2281d1f51fe137eafb914106524"
)

type MicropayController struct {
	controller
}

func (this *MicropayController) Get() {
	this.Data["IsMicroPay"] = true
	this.TplName = "micropay.tpl"
}

func (this *MicropayController) Post() {
	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	amt := this.Input().Get("amt")
	rmbAmt := this.Input().Get("rmbAmt")
	reference := this.Input().Get("reference")
	description := this.Input().Get("description")
	note := this.Input().Get("note")
	ipnUrl := this.Input().Get("ipnUrl")
	token := this.Input().Get("token")
	openid := this.Input().Get("openid")

	if "" == reference {
		reference = fmt.Sprintf("seq_%d", time.Now().Unix())
	}
	if "" == token {
		token = miropayToken
	}
	req := &yuan.Micropay{
		MerchantNo:  merchantNo,
		StoreNo:     storeNo,
		Currency:    "USD",
		Amount:      amt,
		RmbAmount:   rmbAmt,
		Vendor:      "wechatpay",
		Reference:   reference,
		IpnUrl:      ipnUrl,
		Description: description,
		Note:        note,
		Openid:      openid,
	}

	resp, err := req.PostToYuansfer(token)
	if err != nil {
		log.Println(err)
		this.Ctx.WriteString("something wrong happened")
	}
	t := template.New("secury pay template")
	t, _ = t.Parse(resp)
	_ = t.Execute(this.Ctx.ResponseWriter, resp)

	return
}
