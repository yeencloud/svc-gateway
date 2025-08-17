package main

import (
	"context"

	"github.com/yeencloud/bpt-service/internal/adapters/http"
	"github.com/yeencloud/bpt-service/internal/service"
	baseservice "github.com/yeencloud/lib-base"
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
