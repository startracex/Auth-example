package handler

import (
	"github.com/startracex/Auth-example/conf"
	"github.com/startracex/Auth-example/db/query"
	intf "github.com/startracex/Auth-example/interface"
	"github.com/startracex/Auth-example/utils/jwt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	input := c.MustGet("input").(intf.SignInput)
	data := query.QueryUser(query.A{"email": input.Email, "password": input.Password})
	if data == nil {
		c.JSON(400, gin.H{"error": "email or password error"})
		return
	}
	claims := jwt.JWT(data)
	// json web token
	token, err := jwt.Sign256(claims, conf.Global.Webkey)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("Token", token, 3600, "/", "", false, false)
	c.JSON(200, gin.H{"token": token})
}
