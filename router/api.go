package router

import (
	"github.com/gin-gonic/gin"
	"github.com/startracex/Auth-example/handler/api"
	mw "github.com/startracex/Auth-example/middleware"
)

func ApiRouter(e *gin.Engine) {
	g := e.Group("/api")
	g.OPTIONS("/user", mw.CorsOptions)
	g.OPTIONS("/auth", mw.CorsOptions)

	g.GET("/user", mw.CorsAll, api.UserApi)
	g.GET("/auth", mw.CorsAll, api.AuthApi)
}
