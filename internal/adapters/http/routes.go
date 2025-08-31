package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yeencloud/svc-gateway/contract"
	"github.com/yeencloud/svc-gateway/graph"
)

func (s *HTTPServer) registerRoutes(engine *gin.Engine) {
	debug := engine.Group("/debug")
	r := engine.Use(s.server.RequireCorrelationID, s.server.RequireRequestID)

	contract.RegisterRoutes(r, s)

	r.POST("/query", s.graphqlHandler())
	debug.GET("/", s.playgroundHandler())
}

func (s *HTTPServer) graphqlHandler() gin.HandlerFunc {
	h := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (s *HTTPServer) playgroundHandler() gin.HandlerFunc {
	debugUuid := uuid.NewString()

	h := playground.Handler("GraphQL", "/query", playground.WithGraphiqlFetcherHeaders(map[string]string{
		"X-Request-ID":     debugUuid,
		"X-Correlation-ID": debugUuid,
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
