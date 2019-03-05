package controllers

import (
	"encoding/json"
	"log"

	yuan "github.com/yuansfer/golang_sdk"
)

type InstoreReverseController struct {
	Controller
}

func (this *InstoreReverseController) Get() {
	this.Data["IsReverse"] = true
	this.TplName = "instore-reverse.tpl"
}

func (this *InstoreReverseController) Post() {

	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")

	transactionNo := this.Input().Get("transactionNo")

	req := yuan.Reverse{
		MerchantNo:    merchantNo,
		StoreNo:       storeNo,
		TransactionNo: transactionNo,
	}

	ret, err := req.PostToYuansfer()
	if err != nil {
		log.Println(err)
	}
	log.Println(ret)
	var qr yuan.QryReverse

	err = json.Unmarshal([]byte(ret), &qr)
	if err != nil {
		log.Println("Unmarshal Err:%v", err)
	}
	log.Println("resp:", qr)
	resp := qr.Result
	this.Data["IsInquire"] = true
	this.TplName = "inquire.tpl"

	this.checkData("Reference", resp.TransactionNo)
	this.checkData("YuansferId", resp.TransactionStatus)
	this.checkData("Amount", resp.Amount)
	this.checkData("Status", "")
	this.checkData("RefundAmount", resp.RefundAmount)

	return
}
