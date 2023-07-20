package handler

import (
	"fmt"
	"main/conf"
	"main/db/query"
	"main/interface"
	"main/utils/jwt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	input := c.MustGet("input").(intf.SignInput)
	fmt.Println("Mustget input:", input)
	data := query.QueryUser(query.A{"email": input.Email, "password": input.Password})
	if data == nil {
		c.JSON(400, gin.H{"error": "email or password error"})
		return
	}
	claims := jwt.Map(data)
	// json web token
	token, err := jwt.Sign256(claims, conf.Global.Webkey)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("Token", token, 3600, "/", "", false, false)
	c.JSON(200, gin.H{"token": token})
}
