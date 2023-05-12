package router

import (
	"github.com/gin-gonic/gin"
	"main/api"
)

func ApiRouter(e *gin.Engine) {
	g := e.Group("/api")
	g.OPTIONS("/user", CorsOptions)
	g.OPTIONS("/auth", CorsOptions)

	g.GET("/user", Cors, api.UserApi)
	g.GET("/auth", Cors, api.AuthApi)
}
