package service

import (
	"github.com/yeencloud/svc-gateway/internal/ports"
)

type service struct {
	ports *ports.Ports
}

func NewUsecases() service {
	return service{
		ports: &ports.Ports{},
	}
}
