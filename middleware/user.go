package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header
		for k, _ := range headers {
			if strings.HasPrefix(k, "X-Jwt-Claim") || strings.HasPrefix(k, "x-jwt-claim") {
				c.Set(strings.ToLower(k), c.GetHeader(k))
			}
		}
		c.Next()
	}
}
