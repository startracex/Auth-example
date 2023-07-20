package api

import (
	"github.com/gin-gonic/gin"
	"main/db"
)

func UserApi(c *gin.Context) {
	// Need header: Token
	token := c.GetHeader("Token")
	if token == "" {
		c.AbortWithStatus(400)
		return
	}
	strid, err := db.TokenMap.Get(token)
	if err != nil || strid == "" {
		c.AbortWithStatus(400)
		return
	}
	// TODO DANGEROUS DATA DID NOT FILTER
	user := db.SignDB.FindByID(strid)
	c.JSON(200, user)
}
