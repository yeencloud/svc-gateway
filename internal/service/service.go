package service

import (
	"github.com/yeencloud/bpt-service/internal/ports"
)

type service struct {
	ports *ports.Ports
}

func NewUsecases() service {
	return service{
		ports: &ports.Ports{},
	}
}
