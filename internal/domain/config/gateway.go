package config

type GatewayConfig struct {
	URL     string `config:"GATEWAY_URL" default:"localhost:8080"`
	TLS     bool   `config:"GATEWAY_TLS" default:"false"`
	TLSCert string `config:"GATEWAY_TLS_CERT" default:""`
	TLSKey  string `config:"GATEWAY_TLS_KEY" default:""`
}
