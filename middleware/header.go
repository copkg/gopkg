package middleware

import (
	"strings"

	"github.com/copkg/gopkg/context"
	"github.com/gin-gonic/gin"
)

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header
		cc := context.Context{
			Context: c,
		}
		for k, _ := range headers {
			if strings.HasPrefix(strings.ToLower(k), "x-") {
				cc.Set(k, c.GetHeader(k))
			}
		}
		c.Next()
	}
}
