package http

import (
	httpserver "github.com/yeencloud/lib-httpserver"
	"github.com/yeencloud/svc-gateway/internal/ports"
)

type HTTPServer struct {
	server *httpserver.HttpServer

	usecases ports.Usecases
}

func NewHTTPServer(server *httpserver.HttpServer, usecases ports.Usecases) *HTTPServer {
	httpServer := &HTTPServer{
		server:   server,
		usecases: usecases,
	}

	httpServer.registerRoutes(server.Gin)

	return httpServer
}
