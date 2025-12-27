package http

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	graphRuntime "github.com/yeencloud/svc-gateway/contract/graphql/generated"
)

func (s *HTTPServer) registerRoutes(engine *gin.Engine) {
	engine.Static("/s", "./static")

	// Setting up oauth routes
	oauth := engine.Group("/oauth")
	oauth.Match([]string{"POST", "GET"}, "/authorize", s.authorizeClient())
	oauth.POST("/token", s.handleToken())

	// Setting up graphql playground
	debug := engine.Group("/debug")
	debug.GET("/", s.playgroundHandler())

	// Setting up graphql endpoint
	r := engine.Use(s.server.RequireCorrelationID, s.server.RequireRequestID)
	r.POST("/query", s.graphqlHandler())
}

func (s *HTTPServer) graphqlHandler() gin.HandlerFunc {
	config := graphRuntime.Config{
		Resolvers: s.resolvers,
	}

	executableSchema := graphRuntime.NewExecutableSchema(config)

	h := handler.NewDefaultServer(executableSchema)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
