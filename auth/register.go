package auth

import (
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
	to := authbind.ToAuth()
	if to.UserExist() {
		c.Status(409)
		c.Abort()
		return
	} else {
		to.UserRegister()
	}
	c.Status(200)
	c.Abort()
}

func PageRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
