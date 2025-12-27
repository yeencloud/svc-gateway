package http

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
