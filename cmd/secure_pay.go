package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"

	"github.com/yuansfer/golang_sdk"
)

const (
	//ConfigFile is the default configuration file name.
	ConfigFile = "./cfg.toml"
)

func init() {
	var (
		configFile string
		cfg        yuansfer.Configuration
	)

	flag.StringVar(&configFile, "yuansfer", ConfigFile, "config file name ")
	flag.Parse()
	println(configFile)
	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		log.Printf("reading configuration failed: %s\n", err.Error())
	}
	yuansfer.LoadConfiguration(cfg)
}

func main() {
	//securePay()
	//securePayYIP()
	//refund()
	retrieve()
}

func securePay() {
	req := &yuansfer.SecurePay{
		Currency:       "USD",
		SettleCurrency: "USD",
		Amount:         "1.99",
		Vendor:         "alipay",
		Reference:      fmt.Sprintf("%s-%d", "alipay", time.Now().Unix()),
		Terminal:       "ONLINE",
		IpnURL:         "https://t.ldf.fit/callback",
		CallbackURL:    "https://t.ldf.fit/callback",
		Note:           "123",
		Description:    "456",
		//Timeout:        "",
		//GoodsInfo:      "",
		//CreditType:     "",
		//CustomerNo:     "",
		//OSType:         "",
	}

	buf, err := req.PostToYuansfer()
	if nil != err {
		fmt.Println("failed: ", err.Error())
	} else {
		fmt.Println("response ==> ", string(buf))
	}
}

func securePayYIP() {
	req := &yuansfer.SecurePay{
		Currency:       "USD",
		SettleCurrency: "USD",
		Amount:         "1.99",
		Vendor:         "paypal",
		Reference:      fmt.Sprintf("%s-%d", "alipay", time.Now().Unix()),
		Terminal:       "YIP",
		IpnURL:         "https://t.ldf.fit/callback",
		CallbackURL:    "https://t.ldf.fit/callback",
		Note:           "123",
		Description:    "456",
	}

	buf, err := req.PostToYuansfer()
	if nil != err {
		fmt.Println("failed: ", err.Error())
	} else {
		fmt.Println("response ==> ", string(buf))
	}
}

func refund() {
	req := yuansfer.Refund{
		Amount:          "0.01",
		Currency:        "USD",
		Reference:       "alipay-1629771459",
		TransactionNo:   "",
		RefundReference: fmt.Sprintf("refund-%d", time.Now().Unix()),
		SettleCurrency:  "USD",
	}

	buf, err := req.PostToYuansfer()
	if nil != err {
		println("failed: ", err.Error())
	} else {
		println("result: ", string(buf))
	}
}

func retrieve() {
	req := yuansfer.Query{
		Reference:     "alipay-1629771459",
		TransactionNo: "",
	}

	buf, err := req.PostToYuansfer()
	if nil != err {
		println(err.Error())
	} else {
		println(string(buf))
	}
}
