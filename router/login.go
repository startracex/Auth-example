package router

import (
	"github.com/gin-gonic/gin"
	"main/auth"
)

func LoginRouter(g *gin.Engine) {
	g.GET("/login", auth.PageLogin)
	g.POST("/login", auth.Login)
	g.PUT("/login", auth.Login)
}
