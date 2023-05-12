package auth

import (
	"github.com/gin-gonic/gin"
	"main/db"
	"main/utils"
)

func NewAuth(c *gin.Context) {
	var authbind BindAuth
	err := c.ShouldBind(&authbind)
	if authbind == (BindAuth{}) {
		err = c.ShouldBindJSON(&authbind)
	}
	if err != nil {
		c.Status(400)
		c.Abort()
		return
	}
	toauth := authbind.ToAuth()
	id := toauth.GetID()
	if id == "" {
		c.Status(400)
		c.Abort()
		return
	}
	code := utils.GenUUID()
	token := utils.Sha1(utils.GenUUID() + id)
	db.CodeMap.Set(code, token, 300)
	db.TokenMap.Set(token, id, 86400)
	c.JSON(200, gin.H{
		"code": code,
	})
}

func PageAuth(c *gin.Context) {
	c.HTML(200, "auth.html", nil)
}
