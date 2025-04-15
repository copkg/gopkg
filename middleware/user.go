package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header
		for k, _ := range headers {
			if strings.HasPrefix(strings.ToLower(k), "x-jwt-claim") {
				c.Set(strings.ToLower(k), c.GetHeader(k))
			}
		}
		c.Next()
	}
}
