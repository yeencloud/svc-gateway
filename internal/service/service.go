package service

import (
	"context"

	"github.com/yeencloud/svc-gateway/internal/adapters/clients"
	"github.com/yeencloud/svc-gateway/internal/domain/config"
	"github.com/yeencloud/svc-gateway/internal/ports"
	contract "github.com/yeencloud/svc-identity/contract/proto/generated"
)

type service struct {
	ports *ports.Ports

	endpoints *config.EndpointsConfig
	identity  *clients.IdentityClient
}

func NewUsecases(endpointsConfig *config.EndpointsConfig) service {
	identityService := clients.NewIdentityClient(endpointsConfig.IdentityService)

	return service{
		ports:     &ports.Ports{},
		endpoints: endpointsConfig,
		identity:  identityService,
	}
}

func (s service) RegisterUser(ctx context.Context) error {
	r, err := s.identity.Register(ctx, &contract.RegisterObject{
		Email:    "A",
		Password: "B",
	})

	if err != nil {
		return err
	}

	_ = r

	return nil
}
