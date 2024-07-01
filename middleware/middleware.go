package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": http.StatusMethodNotAllowed,
			"msg":  http.StatusText(http.StatusMethodNotAllowed),
			"time": time.Now().Unix(),
		})
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg":  http.StatusText(http.StatusNotFound),
			"time": time.Now().Unix(),
		})
	}
}
