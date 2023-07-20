package mw

import (
	"main/interface"
	"github.com/gin-gonic/gin"
)

// Verify validity, set "input" in context
func VerifyValidity(c *ctx) {
	input, err := intf.ContextSignInput(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if input.Email == "" || input.Password == "" {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.Set("input", input)
}
