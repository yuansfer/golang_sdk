package controllers

import (
	"io"
)

type CallbackController struct {
	Controller
}

func (this *CallbackController) Get() {
	reference := this.Input().Get("reference")
	if "" == reference {
		this.Redirect("/", 302)
	} else {
		this.Data["IsPay"] = true
		this.TplName = "callback.tpl"

		this.checkData("Amt", this.Input().Get("amount"))
		this.checkData("Vendor", this.Input().Get("vendor"))
		this.checkData("RmbAmt", this.Input().Get("rmbAmount"))
		this.checkData("Note", this.Input().Get("note"))
		this.checkData("Reference", this.Input().Get("reference"))

	}
}

func (this *CallbackController) Post() {
	str := "success"
	io.WriteString(this.Ctx.ResponseWriter, str)
}
