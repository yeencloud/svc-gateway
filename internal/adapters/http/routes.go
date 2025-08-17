package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yeencloud/svc-gateway/graph"
)

func (s *HTTPServer) registerRoutes(engine *gin.Engine) {
	debug := engine.Group("/debug")
	r := engine.Use(s.server.RequireCorrelationID, s.server.RequireRequestID)

	r.POST("/query", s.graphqlHandler())
	debug.GET("/", s.playgroundHandler())
}

// Defining the Graphql handler
func (s *HTTPServer) graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
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
