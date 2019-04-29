package controllers

import (
	"encoding/json"
	"log"

	yuan "github.com/yuansfer/golang_sdk"
)

type InquireController struct {
	Controller
}

func (this *InquireController) Get() {
	this.Data["IsInquire"] = true
	this.TplName = "submit-inquire.tpl"
}

func (this *InquireController) Post() {

	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	reference := this.Input().Get("reference")

	req := yuan.Query{
		MerchantNo: merchantNo,
		StoreNo:    storeNo,
		Reference:  reference,
	}

	ret, err := req.PostToYuansfer()
	if err != nil {
		log.Println(err)
	}
	log.Println(ret)
	var qr yuan.QryResponse

	err = json.Unmarshal([]byte(ret), &qr)
	if err != nil {
		log.Println("Unmarshal Err:%v", err)
	}
	log.Println("resp:", qr)
	resp := qr.Result
	this.Data["IsInquire"] = true
	this.TplName = "inquire.tpl"

	this.checkData("Reference", resp.Reference)
	this.checkData("YuansferId", resp.YuansferID)
	this.checkData("Amount", resp.Amount)
	this.checkData("Status", resp.Status)
	this.checkData("RefundInfo", resp.RefundInfo)
	this.checkData("Currency", resp.Currency)

	return
}
