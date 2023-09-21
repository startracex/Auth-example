package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/startracex/Auth-example/db"
	"github.com/startracex/Auth-example/db/query"
	intf "github.com/startracex/Auth-example/interface"
)

func Register(c *gin.Context) {
	input, _ := c.MustGet("input").(intf.SignInput)
	if input.Email == "" || input.Password == "" {
		c.JSON(400, gin.H{"error": "email or password error"})
		return
	}
	exist := query.UserExist(query.A{"email": input.Email})
	if exist {
		c.JSON(400, gin.H{"error": "email already exist"})
		return
	}
	db.SignDB.Add(query.A{"enail": input.Email, "password": input.Password})
	c.String(200, "register success")
}
