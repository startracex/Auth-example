package router

import (
	"github.com/gin-gonic/gin"
	"github.com/startracex/Auth-example/handler"
	mw "github.com/startracex/Auth-example/middleware"
)

func registerEntrance(g *gin.Engine) {
	g.OPTIONS("/register", mw.CorsOptions)
	g.POST("/register", mw.VerifyValidity, handler.Register)
}
