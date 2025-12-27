package http

import (
	httpserver "github.com/yeencloud/lib-httpserver"
	"github.com/yeencloud/svc-gateway/internal/adapters/graphql"
)

type HTTPServer struct {
	server *httpserver.HttpServer

	resolvers *graphql.Resolver
}

func NewHTTPServer(server *httpserver.HttpServer, usecases *graphql.Resolver) *HTTPServer {
	httpServer := &HTTPServer{
		server:    server,
		resolvers: usecases,
	}

	httpServer.registerRoutes(server.Gin)

	return httpServer
}
