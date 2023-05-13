package auth

import (
	"main/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
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
	toauth.Password = utils.Sha1(toauth.Password)
	if toauth.UserExist() {
		c.Status(409)
		c.Abort()
		return
	} else {
		toauth.UserRegister()
	}
	c.Status(200)
	c.Abort()
}

func PageRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
