package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	yuan "github.com/yuansfer/golang_sdk"
)

type InstoreQrcodeController struct {
	Controller
}

func (this *InstoreQrcodeController) Get() {
	this.Data["IsCreateQrcode"] = true
	this.TplName = "instore-create-qrcode.tpl"
}

func (this *InstoreQrcodeController) Post() {

	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")

	amt := this.Input().Get("amt")

	vendor := this.Input().Get("vendor")
	reference := this.Input().Get("reference")
	if "" == reference {
		reference = fmt.Sprintf("seq_%d", time.Now().Unix())
	}
	ipnUrl := this.Input().Get("ipnUrl")

	req := yuan.CreateQrcode{
		MerchantNo: merchantNo,
		StoreNo:    storeNo,
		Amount:     amt,
		Vendor:     vendor,
		Reference:  reference,
		IpnUrl:     ipnUrl,
		Currency:   "USD",
	}

	ret, err := req.PostToYuansfer()
	if err != nil {
		log.Println(err)
	}
	log.Println(ret)
	var qr yuan.QrcodeRet
	err = json.Unmarshal([]byte(ret), &qr)
	if err != nil {
		log.Println("Unmarshal Err:%v", err)
	}
	log.Println("resp:", qr)
	resp := qr
	this.Data["IsCreateQrcode"] = true
	this.TplName = "instore-create-qrcode-ret.tpl"

	this.checkData("RectCode", resp.RectCode)
	this.checkData("DeepLink", resp.DeepLink)
	this.checkData("TransactionNo", resp.TransactionNo)
	this.checkData("QrcodeUrl", resp.QrcodeUrl)
	this.checkData("TimeOut", resp.TimeOut)
	this.checkData("Reference", resp.Reference)

	return
}
