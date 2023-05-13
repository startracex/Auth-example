package router

import (
	"github.com/gin-gonic/gin"
	"main/api"
)

func ApiRouter(e *gin.Engine) {
	g := e.Group("/api")
	g.OPTIONS("/user", CorsOptions)
	g.OPTIONS("/auth", CorsOptions)

	g.GET("/user", Cors, api.UserApi)
	g.GET("/auth", Cors, api.AuthApi)
}

func Cors(c *gin.Context) {
	CorsContext(c)
	c.Next()
}

func CorsOptions(c *gin.Context) {
	CorsContext(c)
}

func CorsContext(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
}
