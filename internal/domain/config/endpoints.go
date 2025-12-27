package config

type EndpointsConfig struct {
	IdentityService string `config:"SVC_IDENTITY_URL" default:"localhost:8081"`
}
