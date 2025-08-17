package graph

import (
	"context"

	"github.com/samber/lo"
	"github.com/yeencloud/svc-gateway/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r *Resolver) getProfile(ctx context.Context, id string) (*model.Profile, error) {
	_ = ctx

	return &model.Profile{
		ID:  id,
		Bio: lo.ToPtr("Tiefling Archer Lv. 12"),
	}, nil
}
