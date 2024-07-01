package middleware

import (
	"github.com/copkg/gopkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"time"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				logger.WithFields(map[string]interface{}{
					"time":    time.Now(),
					"error":   err,
					"path":    ctx.Request.URL.Path,
					"request": string(httpRequest),
					"stack":   string(debug.Stack()),
				}).Error("[Recovery from panic]")
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  http.StatusText(http.StatusInternalServerError),
					"time": time.Now().Unix(),
				})
			}
		}()
		ctx.Next()
	}
}
