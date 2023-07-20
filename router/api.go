package router

import (
	"github.com/gin-gonic/gin"
	"main/api"
	"main/middleware"
)

func ApiRouter(e *gin.Engine) {
	g := e.Group("/api")
	g.OPTIONS("/user", mw.CorsOptions)
	g.OPTIONS("/auth", mw.CorsOptions)

	g.GET("/user", mw.CorsAll, api.UserApi)
	g.GET("/auth", mw.CorsAll, api.AuthApi)
}
