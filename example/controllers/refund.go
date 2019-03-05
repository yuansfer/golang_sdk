package controllers

import (
	// "html/template"
	"encoding/json"
	"log"

	yuan "github.com/yuansfer/golang_sdk"
)

type RefundController struct {
	Controller
}

func (this *RefundController) Get() {
	this.Data["IsRefund"] = true
	this.TplName = "submit-refund.tpl"
}

func (this *RefundController) Post() {

	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	reference := this.Input().Get("reference")
	password := this.Input().Get("password")
	amt := this.Input().Get("amt")
	rmbAmt := this.Input().Get("rmbAmt")
	managerAccountNo := this.Input().Get("managerAccountNo")

	req := yuan.Refund{
		MerchantNo:       merchantNo,
		StoreNo:          storeNo,
		Reference:        reference,
		Password:         password,
		Amount:           amt,
		RmbAmount:        rmbAmt,
		ManagerAccountNo: managerAccountNo,
	}

	ret, err := req.PostToYuansfer()
	if err != nil {
		log.Println(err)
		this.Ctx.WriteString("something wrong happened in refund")

		return
	}

	var qr yuan.RefundResponse

	err = json.Unmarshal([]byte(ret), &qr)
	if err != nil {
		log.Println("Unmarshal Err:%v", err)
	}
	log.Println("resp:", qr)
	resp := qr.Result

	this.checkData("Reference", resp.Reference)
	this.checkData("YuansferId", resp.OldTransactionId)
	this.checkData("Amount", resp.RefundTransactionId)
	this.checkData("Status", resp.Status)

	this.Data["IsRefund"] = true

	if "000100" == qr.RetCode && "success" == resp.Status {
		this.TplName = "refund-success.tpl"
	} else {
		this.Ctx.WriteString("refund failed")
	}

	return
}
