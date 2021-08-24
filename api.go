package yuansfer

// Yuansfer payment method
type Yuansfer interface {
	PostToYuansfer() ([]byte, error)
}

const (
	URISecurePay    = "/online/v3/secure-pay"
	URIYipApp       = "/micropay/v3/prepay"
	URIYipProcess   = "/creditpay/v3/process"
	URIRefund       = "/app-data-search/v3/refund"
	URICancel       = "/app-data-search/v3/cancel"
	URIRetrieve     = "/app-data-search/v3/tran-query"
	URIPosAdd       = "/app-instore/v3/add"
	URIPosPay       = "/app-instore/v3/pay"
	URIPosCreateQRC = "/app-instore/v3/create-trans-qrcode"

	GatewaySandbox    = "https://mapi.yuansfer.yunkeguan.com"
	GatewayProduction = "https://mapi.yuansfer.com"
)
