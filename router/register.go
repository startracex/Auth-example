package router

import (
	"main/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(g *gin.Engine) {
	g.GET("/register", auth.PageRegister)
	g.POST("/register", auth.Register)
	g.PUT("/register", auth.Register)
}
