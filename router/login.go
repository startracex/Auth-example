package router

import (
	"main/handler"
	"main/middleware"
	"github.com/gin-gonic/gin"
)

func loginEntrance(g *gin.Engine) {
	g.OPTIONS("/login", mw.CorsOptions)
	g.POST("/login", mw.VerifyValidity, handler.Login)
}
