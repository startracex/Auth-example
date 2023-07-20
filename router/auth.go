package router

import (
	"main/handler"
	mw "main/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRouter(g *gin.Engine) {
	g.OPTIONS("/auth", mw.CorsOptions)
	g.POST("/auth", handler.NewAuth)
}
