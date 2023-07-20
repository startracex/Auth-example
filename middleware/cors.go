package mw

import (
	"main/conf"
)

// Allow cors: allowed in config
func CorsSome(c *ctx) {
	origin := c.Request.Header.Get("Origin")
	for _, v := range conf.Global.Allow {
		if v == origin {
			CorsContext(c, origin)
		}
	}
	c.Next()
}

// Allow cors: *
func CorsAll(c *ctx) {
	CorsContext(c, "*")
	c.Next()
}

// Cors for OPTIONS method
func CorsOptions(c *ctx) {
	CorsContext(c, "*")
	c.AbortWithStatus(200)
}

// Allow cors: allowed
func CorsContext(c *ctx, allowed string) {
	c.Header("Access-Control-Allow-Origin", allowed)
	c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
}
