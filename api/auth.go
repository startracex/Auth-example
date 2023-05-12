package api

import (
	"github.com/gin-gonic/gin"
	"main/db"
)

func AuthApi(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.Status(400)
		c.Abort()
		return
	}
	token, _ := db.CodeMap.Get(code)
	db.CodeMap.Del(code)
	if token == "" {
		c.Status(400)
		c.Abort()
		return
	} else {
		c.JSON(200, gin.H{
			"token": token,
		})
	}
}
