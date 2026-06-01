package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error":             http.StatusText(http.StatusMethodNotAllowed),
			"error_description": "Method Not Allowed",
			"time":              time.Now().Unix(),
		})
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":             http.StatusText(http.StatusNotFound),
			"error_description": "the path is missing",
			"time":              time.Now().Unix(),
		})
	}
}
