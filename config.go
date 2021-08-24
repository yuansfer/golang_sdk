package yuansfer

import (
	"time"
)

var (
	//Development or Production domain
	APIGateway string

	// token obtained from Yuansfer portal
	Token string
	//YuansferAPI is the configuration information
	cfg Configuration

	Cred = make(map[string]string)
)

type Configuration struct {
	Credential Credential `toml:"credentials"`
	Network    Network    `toml:"http"`
}

type Credential struct {
	MerchantNo string `toml:"merchant_no"`
	StoreNo    string `toml:"store_no"`
	Token      string `toml:"token"`
	Env        string `toml:"environment"`
	PartnerNo  string `toml:"partner_no"`
}

// Network contains common configuration.
type Network struct {
	RequestTimeout time.Duration `toml:"request_timeout"`
	ConnectTimeout time.Duration `toml:"connect_timeout"`
	SocketTimeout  time.Duration `toml:"socket_timeout"`
}

func LoadConfiguration(c Configuration) {
	cfg = c
	if "" == cfg.Credential.Env {
		cfg.Credential.Env = "sandbox"
	}
	if "production" == cfg.Credential.Env {
		APIGateway = GatewayProduction
	} else {
		APIGateway = GatewaySandbox
	}

	Token = cfg.Credential.Token

	if 0 != cfg.Network.RequestTimeout {
		network.RequestTimeout = cfg.Network.RequestTimeout
	}

	if 0 != cfg.Network.ConnectTimeout {
		network.ConnectTimeout = cfg.Network.ConnectTimeout
	}

	if 0 != cfg.Network.SocketTimeout {
		network.SocketTimeout = cfg.Network.SocketTimeout
	}

	Cred["merchantNo"] = cfg.Credential.MerchantNo
	Cred["merGroupNo"] = cfg.Credential.PartnerNo
	Cred["storeNo"] = cfg.Credential.StoreNo
}
