package router

import (
	"github.com/gin-gonic/gin"
	"main/handler"
	"main/middleware"
)

func registerEntrance(g *gin.Engine) {
	g.OPTIONS("/register", mw.CorsOptions)
	g.POST("/register", mw.VerifyValidity, handler.Register)
}
