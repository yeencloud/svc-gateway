package contract

import "github.com/gin-gonic/gin"

type RoutesHandler interface {
	RegisterUser(context *gin.Context)
}

func RegisterRoutes(gin gin.IRoutes, handler RoutesHandler) {
	gin.POST("/users", handler.RegisterUser)
}
