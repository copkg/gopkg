package response

import (
	"github.com/copkg/gopkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	OKStatus    int    = http.StatusOK
	ErrorStatus string = "ERROR"
	FailStatus  int    = 0
)

type ResponseResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, v interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, ResponseResult{
		Code:    OKStatus,
		Message: "success",
		Data:    v,
	})
}
func Fail(c *gin.Context, err error, status ...int) {
	code := http.StatusBadRequest
	if len(status) > 0 {
		code = status[0]
	}
	if code >= 500 {
		logger.Errorf("Internal server error", err)
	}
	c.AbortWithStatusJSON(code, ResponseResult{
		Code:    FailStatus,
		Message: err.Error(),
	})
}
