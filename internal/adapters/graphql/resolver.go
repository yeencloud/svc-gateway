package graphql

import (
	"context"

	"github.com/samber/lo"
	"github.com/yeencloud/svc-gateway/contract/graphql/generated"
	"github.com/yeencloud/svc-gateway/internal/ports"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	Usecases ports.Usecases
}

func (r *Resolver) getProfile(ctx context.Context, id string) (*generated.Profile, error) {
	_ = ctx

	return &generated.Profile{
		ID:  id,
		Bio: lo.ToPtr("Tiefling Archer Lv. 12"),
	}, nil
}
