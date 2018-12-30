package yuansfer

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type yuansferApi struct {
	OnlineHost          []string `yaml:"online_host"`
	OnlinePayment       string   `yaml:"online_payment_url"`
	OnlineQuery         string   `yaml:"online_query_url"`
	OnlineRefund        string   `yaml:"online_refund_url"`
	InstoreHost         []string `yaml:"instore_host"`
	InstoreAdd          string   `yaml:"instore_add_url"`
	InstorePay          string   `yaml:"instore_pay_url"`
	InstoreCreateQrcode string   `yaml:"instore_create_qrcode"`
	InstoreQuery        string   `yaml:"instore_query_url"`
	InstoreRefund       string   `yaml:"instore_refund_url"`
	InstoreReverse      string   `yaml:"instore_reverse_url"`
}

const (
	CONFIG_FILE = "./yuansfer-api.yaml"
)

func init() {

	env := flag.String("env", "dev", "enviroment dev or product")
	flag.Parse()

	data, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		log.Fatalf("read config file %s error:%v", CONFIG_FILE, err)
		os.Exit(21)
	}
	err = yaml.Unmarshal([]byte(data), &YuansferApi)
	if err != nil {
		log.Fatalf("unmarshal config file %s error:%v", CONFIG_FILE, err)
		os.Exit(22)
	}

	if "product" == *env {
		yuansferHost = YuansferApi.OnlineHost[1]
		instoreHost = YuansferApi.InstoreHost[1]
	} else {
		yuansferHost = YuansferApi.OnlineHost[0]
		instoreHost = YuansferApi.InstoreHost[0]
	}

}
