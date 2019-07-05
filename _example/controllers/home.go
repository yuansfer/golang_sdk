package controllers

import (
	"fmt"
	"html/template"
	"time"

	yuan "github.com/yuansfer/golang_sdk"
)

type HomeController struct {
	Controller
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
	ipnURL := this.Input().Get("ipnUrl")
	callbackURL := this.Input().Get("callbackUrl")
	creditType := this.Input().Get("creditType")
	paymentCount := this.Input().Get("paymentCount")
	frequency := this.Input().Get("frequency")
	customerNo := ""

	if "" == terminal {
		terminal = "ONLINE"
	}
	if "" == reference {
		reference = fmt.Sprintf("seq_%d", time.Now().Unix())
	}

	req := &yuan.Securepay{
		MerchantNo:   merchantNo,
		StoreNo:      storeNo,
		Currency:     "USD",
		Amount:       amt,
		RmbAmount:    rmbAmt,
		Vendor:       vendor,
		Reference:    reference,
		IpnURL:       ipnURL,
		CallbackURL:  callbackURL,
		Description:  description,
		Note:         note,
		Terminal:     terminal,
		Timeout:      "15",
		CreditType:   creditType,
		PaymentCount: paymentCount,
		Frequency:    frequency,
		CustomerNo:   customerNo,
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

	resp, err := req.PostToYuansfer()
	if err != nil {
		fmt.Println(err)
		this.Ctx.WriteString("something wrong happened")
	}
	t := template.New("secure pay template")
	t, _ = t.Parse(resp)
	_ = t.Execute(this.Ctx.ResponseWriter, resp)

	return
}
