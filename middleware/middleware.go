package middleware

import (
	"github.com/copkg/gopkg/context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() context.HandlerFunc {
	return func(c *context.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    http.StatusMethodNotAllowed,
			"message": http.StatusText(http.StatusMethodNotAllowed),
			"time":    time.Now().Unix(),
		})
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() context.HandlerFunc {
	return func(c *context.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": http.StatusText(http.StatusNotFound),
			"time":    time.Now().Unix(),
		})
	}
}
