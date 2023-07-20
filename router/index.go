package router

import (
	"github.com/gin-gonic/gin"
	"main/conf"
	"main/middleware"
	"main/utils/fs"
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
