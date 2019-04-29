![Yuansfer](http://oss.yuansfer.com/log_20190410.png?x-oss-process=image/resize,l_300)

# golang_sdk [![Build Status](https://travis-ci.org/yuansfer/golang_sdk.svg?branch=master)](https://travis-ci.org/yuansfer/golang_sdk) [![GoDoc](https://godoc.org/github.com/yuansfer/golang_sdk?status.svg)](https://godoc.org/github.com/yuansfer/golang_sdk)
yuansfer SDK for golang language

## about
Add a configuration file called `config.toml` into your project. You can also use the command-line flag `conf` to use a configuration file with a special name.
```
yuansfer_host:
  # sandbox host url
  - https://mapi.yuansfer.yunkeguan.com
  # product host url
  - https://mapi.yuansfer.com
online_payment_url: /appTransaction/v2/securepay
online_query_url: /appTransaction/v2/securepay-reference-query
online_refund_url: /appTransaction/v2/securepayRefund
instore_add_url: /app-instore/v2/add
instore_pay_url: /app-instore/v2/pay
instore_create_qrcode: /app-instore/v2/create-trans-qrcode
instore_refund_url: /app-instore/v2/refund
instore_reverse_url: /app-instore/v2/reverse
micropay_url: /micropay/v2/prepay
token:
  # please get Yuansfer Token from support@yuansfer.com
  online_token: 'ONLINE_TOKEN'
  instore_token: 'INSTORE_TOKEN'
  micropay_token: 'MICROPAY_TOKEN'
# please get the value of Password Prefix at:
# https://docs.yuansfer.com/en/#notice
password_prefix: 'PASSWORD_PREFIX'
```

Use the command-line flag `env` to select a Development or Production environment.

## contact us
mail: support@yuansfer.com
API documention: [https://docs.yuansfer.com/en/](https://docs.yuansfer.com/en/)
