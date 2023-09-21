package router

import (
	"github.com/gin-gonic/gin"
	"github.com/startracex/Auth-example/handler"
	mw "github.com/startracex/Auth-example/middleware"
)

func loginEntrance(g *gin.Engine) {
	g.OPTIONS("/login", mw.CorsOptions)
	g.POST("/login", mw.VerifyValidity, handler.Login)
}
