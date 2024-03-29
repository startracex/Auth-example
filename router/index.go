package router

import (
	"github.com/gin-gonic/gin"
	"github.com/startracex/Auth-example/conf"
	mw "github.com/startracex/Auth-example/middleware"
	"github.com/startracex/Auth-example/utils/fs"
)

func Entrance() {
	if conf.Global.Env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	if len(conf.Global.Allow) > 0 {
		g.Use(mw.CorsSome)
	}

	// Load html
	g.LoadHTMLFiles(fs.WalkExt("view", "html")...)
	g.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})
	registerEntrance(g)
	loginEntrance(g)
	ApiRouter(g)
	AuthRouter(g)
	g.NoRoute(NoRoute)

	g.Run(":" + conf.Global.Port)
}

func NoRoute(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.AbortWithStatus(404)
		return
	}

	if c.Request.URL.Path[len(c.Request.URL.Path)-1] == '/' {
		c.Redirect(301, c.Request.URL.Path[:len(c.Request.URL.Path)-1])
		return
	}

	c.Status(500)
}
