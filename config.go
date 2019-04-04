package yuansfer

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type yuansferApi struct {
	Host                []string      `yaml:"yuansfer_host" toml:"yuansfer_host"`
	OnlinePayment       string        `yaml:"online_payment_url" toml:"online_payment_url"`
	OnlineQuery         string        `yaml:"online_query_url" toml:"online_query_url"`
	OnlineRefund        string        `yaml:"online_refund_url" toml:"online_refund_url"`
	InstoreAdd          string        `yaml:"instore_add_url" toml:"instore_add_url"`
	InstorePay          string        `yaml:"instore_pay_url" toml:"instore_pay_url"`
	InstoreCreateQrcode string        `yaml:"instore_create_qrcode" toml:"instore_create_qrcode"`
	InstoreQuery        string        `yaml:"instore_query_url" toml:"instore_query_url"`
	InstoreRefund       string        `yaml:"instore_refund_url" toml:"instore_refund_url"`
	InstoreReverse      string        `yaml:"instore_reverse_url" toml:"instore_reverse_url"`
	Micropay            string        `yaml:"micropay_url" toml:"micropay_url"`
	PwdPre              string        `yaml:"password_prefix" toml:"password_prefix"`
	Token               yuansferToken `yaml:"token" toml:"token"`
}

type yuansferToken struct {
	SecurepayToken string `yaml:"online_token" toml:"online_token"`
	InstoreToken   string `yaml:"instore_token" toml:"instore_token"`
	MicropayToken  string `yaml:"micropay_token" toml:"micropay_token"`
}

const (
	CONFIG_FILE = "./yuansfer-api.yaml"
)

func init() {
	var configFile string

	env := flag.String("env", "dev", "enviroment dev or product ")
	configFile = *flag.String("conf", CONFIG_FILE, "config file name ")
	flag.Parse()

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("read config file %s error:%s", configFile, err.Error())
		os.Exit(21)
	}
	err = yaml.Unmarshal([]byte(data), &YuansferApi)
	if err != nil {
		log.Fatalf("Unmarshal config file %s error:%s", configFile, err.Error())
		os.Exit(22)
	}

	if "product" == *env {
		yuansferHost = YuansferApi.Host[1]
	} else {
		yuansferHost = YuansferApi.Host[0]
	}
}
