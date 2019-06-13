package controllers

import (
	"io"

	yuan "github.com/yuansfer/golang_sdk"
)

type CallbackController struct {
	Controller
}

func (this *CallbackController) Get() {
	token := yuan.YuansferAPI.Token.SecurepayToken
	request := this.Input()
	_, ret := yuan.VerifySignNotify(request, token)

	if false == ret {
		this.Ctx.WriteString("verifySign Rejected")
	}

	this.Data["IsPay"] = true
	this.TplName = "callback.tpl"

	this.checkData("Amt", request.Get("amount"))
	this.checkData("Status", request.Get("status"))
	this.checkData("RmbAmt", request.Get("rmbAmount"))
	this.checkData("Note", request.Get("note"))
	this.checkData("YuansferID", request.Get("yuansferId"))
	this.checkData("Reference", request.Get("reference"))
}

func (this *CallbackController) Post() {
	token := yuan.YuansferAPI.Token.SecurepayToken
	request := this.Input()
	_, ret := yuan.VerifySignNotify(request, token)

	response := ""
	if ret == true {
		response = "success"
	} else {
		response = "failed"
	}
	io.WriteString(this.Ctx.ResponseWriter, response)
}
