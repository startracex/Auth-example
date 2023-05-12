package auth

import (
	"github.com/gin-gonic/gin"
	"main/utils"
)

func Login(c *gin.Context) {
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
	data := toauth.UserLogin()
	utils.SafeInfo(&data)
	if len(data) == 0 {
		c.Status(401)
		c.Abort()
		return
	}
	c.JSON(200, data)
}

func PageLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
