package mw

import (
	"github.com/gin-gonic/gin"
	intf "github.com/startracex/Auth-example/interface"
)

// Verify validity, set "input" in context
func VerifyValidity(c *gin.Context) {
	input, err := intf.ContextSignInput(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.Set("input", input)
}
