package main

import (
	"context"

	baseservice "github.com/yeencloud/lib-base"
	"github.com/yeencloud/svc-gateway/internal/adapters/http"
	"github.com/yeencloud/svc-gateway/internal/service"
)

func main() {
	baseservice.Run("base-service", baseservice.Options{
		UseDatabase: false,
		UseEvents:   false,
	}, func(ctx context.Context, svc *baseservice.BaseService) error {
		httpServer, err := svc.GetHttpServer()
		if err != nil {
			return err
		}

		usecases := service.NewUsecases()
		http.NewHTTPServer(httpServer, usecases)

		return nil
	})
}
