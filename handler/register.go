package handler

import (
	"github.com/gin-gonic/gin"
	"main/db/query"
	"main/interface"
)

func Register(c *gin.Context) {
	input, err := intf.ContextSignInput(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if input.Email == "" || input.Password == "" {
		c.JSON(400, gin.H{"error": "email or password error"})
		return
	}
	exist := query.UserExist(query.A{"email": input.Email})
	if exist {
		c.JSON(400, gin.H{"error": "email already exist"})
		return
	}
	// TODO add user
	c.String(200, "register success")
}
