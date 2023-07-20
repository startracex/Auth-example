package utils

import "github.com/gin-gonic/gin"

func E(c *gin.Context, err error) {
	if err != nil {
		c.AbortWithStatus(400)
	}
}

func S(c *gin.Context, str string) {
	if str == "" {
		c.AbortWithStatus(400)
	}
}
