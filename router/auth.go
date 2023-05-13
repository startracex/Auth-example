package router

import (
	"main/auth"
	"github.com/gin-gonic/gin"
)

func AuthRouter(g *gin.Engine) {
	g.GET("/auth",  auth.PageAuth)
	g.POST("/auth", auth.NewAuth)
	g.PUT("/auth",  auth.NewAuth)
}