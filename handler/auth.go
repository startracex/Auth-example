package handler

import (
	"main/conf"
	"main/db"
	"main/utils"
	"main/utils/jwt"
	"github.com/gin-gonic/gin"
)

// Feedback one-time access code
func NewAuth(c *gin.Context) {

	token := c.GetHeader("Token")
	claims, _, err := jwt.Parse(token, conf.Global.Webkey)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	// Queried ID
	strid := db.ObjectIDToString(claims.Object_ID)

	// Code for api /api/auth?code
	code := utils.GenUUID()

	// Mediation string: value of code, key of token
	medstr := utils.Sha1(utils.GenUUID() + strid)
	db.CodeMap.Set(code, medstr, 300)
	db.TokenMap.Set(medstr, strid, 86400)
	c.JSON(200, gin.H{
		"code": code,
	})
}