package router

import (
	"github.com/gin-gonic/gin"
	"github.com/startracex/Auth-example/handler"
	mw "github.com/startracex/Auth-example/middleware"
)

func AuthRouter(g *gin.Engine) {
	g.OPTIONS("/auth", mw.CorsOptions)
	g.POST("/auth", handler.NewAuth)
}
