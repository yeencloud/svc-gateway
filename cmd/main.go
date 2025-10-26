package main

import (
	"context"

	baseservice "github.com/yeencloud/lib-base"
	"github.com/yeencloud/lib-shared/config"
	"github.com/yeencloud/svc-gateway/internal/adapters/http"
	gatewayConfig "github.com/yeencloud/svc-gateway/internal/domain/config"
	"github.com/yeencloud/svc-gateway/internal/service"
)

func main() {
	baseservice.Run("svc-gateway", baseservice.Options{
		UseDatabase: false,
		UseEvents:   false,
	}, func(ctx context.Context, svc *baseservice.BaseService) error {
		httpServer, err := svc.GetHttpServer()
		if err != nil {
			return err
		}

		endpointsConfig, err := config.FetchConfig[gatewayConfig.EndpointsConfig]()
		if err != nil {
			return err
		}

		usecases := service.NewUsecases(endpointsConfig)
		http.NewHTTPServer(httpServer, usecases)

		return nil
	})
}
