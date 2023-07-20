package mw

import "github.com/gin-gonic/gin"

// 给gin.Context改名为ctx,并共享原有的所有方法
type ctx = gin.Context