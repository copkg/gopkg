package middleware

import (
	"fmt"
	"github.com/copkg/gopkg/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := request.BearerExtractor{}.ExtractToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  err.Error(),
				"time": time.Now().Unix(),
			})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Conf.GetString("jwt.signingkey")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  err.Error(),
				"time": time.Now().Unix(),
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("uid", claims["id"])
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "missing id from token",
				"time": time.Now().Unix(),
			})
			return
		}
		c.Next()
	}
}
