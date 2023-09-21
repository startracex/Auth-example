package api

import (
	"github.com/gin-gonic/gin"
	"github.com/startracex/Auth-example/db"
)

// Feedback token for user id
func AuthApi(c *gin.Context) {
	// Need url query: code
	code := c.Query("code")
	if code == "" {
		c.AbortWithStatus(400)
		return
	}
	// Mediation string: value of code, key of token
	mdstr, err := db.CodeMap.Get(code)
	if err != nil || mdstr == "" {
		c.AbortWithStatus(400)
		return
	}
	// Delete the one-time token
	db.CodeMap.Del(code)
	c.JSON(200, gin.H{
		"token": mdstr,
	})
}
