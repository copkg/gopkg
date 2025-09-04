package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header
		for k, _ := range headers {
			if strings.HasPrefix(strings.ToLower(k), "x-") {
				c.Set(k, c.GetHeader(k))
			}
		}
		c.Next()
	}
}
