package http

import (
	"context"
	"errors"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/yeencloud/lib-shared/apperr"
	"github.com/yeencloud/svc-gateway/internal/domain"
)

func (s *HTTPServer) authorizeClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		if stored, ok := s.server.SessionGet(c, domain.OauthReturnUriSessionKey); ok {
			if formValues, ok := stored.(url.Values); ok {
				c.Request.Form = formValues
			}
			_ = s.server.SessionDelete(c, domain.OauthReturnUriSessionKey)
		} else {
			if err := c.Request.ParseForm(); err != nil {
				err = errors.Join(err, apperr.InvalidArgumentError{})
				s.server.ReplyWithError(c, err)
				return
			}
		}

		c.Request = c.Request.WithContext(
			context.WithValue(c.Request.Context(), "gin_context", c),
		)

		/*if err := s.oauth.HandleAuthorizeRequest(c.Writer, c.Request); err == nil {
			_ = s.server.SessionDelete(c, "LoggedInUserID")
		}*/
	}
}

func (s *HTTPServer) handleToken() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
