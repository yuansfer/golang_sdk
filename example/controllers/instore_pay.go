package controllers

import (
	"encoding/json"
	"log"

	yuan "github.com/yuansfer/golang_sdk/yuansfer"
)

type InstorePayController struct {
	controller
}

func (this *InstorePayController) Get() {
	this.Data["IsInstorePay"] = true
	this.TplName = "instore-pay.tpl"
}

func (this *InstorePayController) Post() {

	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")

	token := this.Input().Get("token")
	transactionNo := this.Input().Get("transactionNo")
	paymentBarcode := this.Input().Get("paymentBarcode")
	vendor := this.Input().Get("vendor")

	if "" == token {
		token = instoreToken
	}

	req := yuan.InstorePay{
		MerchantNo:     merchantNo,
		StoreNo:        storeNo,
		TransactionNo:  transactionNo,
		PaymentBarcode: paymentBarcode,
		Vendor:         vendor,
	}

	ret, err := req.PostToYuansfer(token)
	if err != nil {
		log.Println(err)
	}
	log.Println(ret)
	var qr yuan.PayResponse
	err = json.Unmarshal([]byte(ret), &qr)
	if err != nil {
		log.Println("Unmarshal Err:%v", err)
	}
	log.Println("resp:%v", qr)
	resp := qr.Transaction
	this.Data["IsInstorePay"] = true
	this.TplName = "instore-pay-ret.tpl"

	this.checkData("TransactionNo", resp.TransactionNo)
	this.checkData("OriginalTransactionNo", resp.OriginalTransactionNo)
	this.checkData("Amount", resp.Amount)
	this.checkData("Vendor", resp.Vendor)
	this.checkData("MerchantNo", resp.MerchantNo)
	this.checkData("TransactionStatus", resp.TransactionStatus)

	return
}
