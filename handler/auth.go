package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/startracex/Auth-example/conf"
	"github.com/startracex/Auth-example/db"
	"github.com/startracex/Auth-example/utils"
	"github.com/startracex/Auth-example/utils/jwt"
)

// Feedback one-time access code
func NewAuth(c *gin.Context) {
	token := c.GetHeader("Token")
	jwtToken, err := jwt.Parse(token, conf.Global.Webkey)
	if err != nil {
		c.AbortWithStatus(400)
		return
	} // Queried ID
	strid, ok := jwt.Get(jwtToken.Claims, "_id")
	if !ok {
		c.AbortWithStatus(400)
		return
	}
	// Code for api /api/auth?code
	code := utils.GenUUID()

	// Mediation string: value of code, key of token
	medstr := utils.Sha1(utils.GenUUID())
	db.CodeMap.Set(code, medstr, 300)
	db.TokenMap.Set(medstr, strid.(string), 86400)
	c.JSON(200, gin.H{
		"code": code,
	})
}
