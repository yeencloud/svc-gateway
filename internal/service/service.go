package service

import (
	"github.com/yeencloud/svc-gateway/internal/domain/config"
	"github.com/yeencloud/svc-gateway/internal/ports"
)

type service struct {
	ports *ports.Ports

	endpoints *config.EndpointsConfig
}

func NewUsecases(endpointsConfig *config.EndpointsConfig) service {
	return service{
		ports:     &ports.Ports{},
		endpoints: endpointsConfig,
	}
}
