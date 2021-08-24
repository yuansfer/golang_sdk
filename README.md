
![Yuansfer](http://oss.yuansfer.com/log_20190410.png?x-oss-process=image/resize,l_300)

# golang_sdk [![Build Status](https://travis-ci.org/yuansfer/golang_sdk.svg?branch=master)](https://travis-ci.org/yuansfer/golang_sdk) [![GoDoc](https://godoc.org/github.com/yuansfer/golang_sdk?status.svg)](https://godoc.org/github.com/yuansfer/golang_sdk)
yuansfer SDK for golang 

## about
- Use `yuansfer.LoadConfiguration(y yuansfer.Configuration) function to setup credentials`

```
[credentials]
store_no = "3xxxxx"
merchant_no = "2xxxxx"
environment = "sandbox" # production
token = "obtain your developer token from yuansfer portal"
```

---

```
import "github.com/yuansfer/golang_sdk"
```

---

```go
package main

import "fmt"
import "time"

import "github.com/yuansfer/golang_sdk"

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

```
## contact us
- mail: support@yuansfer.com
- API documentation: [https://docs.yuansfer.com/](https://docs.yuansfer.com/)
