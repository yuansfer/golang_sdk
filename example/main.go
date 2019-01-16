package main

import (
	"github.com/astaxie/beego"
	"github.com/yuansfer/golang_sdk/example/controllers"
)

func main() {
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.Graceful = true

	beego.BConfig.Listen.HTTPSPort = 443
	beego.BConfig.Listen.HTTPSCertFile = "./server.pem"
	beego.BConfig.Listen.HTTPSKeyFile = "./server.key"
	//
	// 注册 beego 路由
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/pay", &controllers.HomeController{})
	beego.Router("/inquire", &controllers.InquireController{})
	// beego.Router("/exchange-rate", &controllers.ExchangeRateController{})

	beego.Router("/callback", &controllers.CallbackController{})

	// In-store
	//
	beego.Router("/instore-add", &controllers.InstoreAddController{})
	beego.Router("/instore-pay", &controllers.InstorePayController{})
	beego.Router("/instore-qrcode", &controllers.InstoreQrcodeController{})
	beego.Router("/instore-reverse", &controllers.InstoreReverseController{})
	beego.Router("/micropay", &controllers.MicropayController{})
	// 启动 beego
	beego.Run()
}
