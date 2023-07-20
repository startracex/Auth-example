package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// BindWithHeader bind data with header, support json, yaml, FormData(default)
func BindWithHeader(c *gin.Context, data any) error {
	header := c.GetHeader("Content-Type")
	switch header {
	case "application/json":
		err := c.ShouldBindWith(&data, binding.JSON)
		return err
	case "application/yaml":
		err := c.ShouldBindWith(&data, binding.YAML)
		return err
	default:
		err := c.ShouldBindWith(&data, binding.Form)
		return err
	}
}
