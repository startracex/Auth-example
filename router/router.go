package router

import (
	"main/config"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
)

func Router() {
	if config.Global.Env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	g.LoadHTMLFiles(WalkHtml()...)

	RegisterRouter(g)
	LoginRouter(g)
	ApiRouter(g)
	AuthRouter(g)
	
	g.Run(":" + config.Global.Port)
}
func WalkHtml() []string {
	var htmls []string
	filepath.Walk("view", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".html" {
			htmls = append(htmls, path)
		}
		return nil
	})
	return htmls
}
