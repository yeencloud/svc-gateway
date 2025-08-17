package http

import (
	"github.com/yeencloud/bpt-service/internal/ports"
	httpserver "github.com/yeencloud/lib-httpserver"
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
