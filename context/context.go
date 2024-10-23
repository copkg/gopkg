package context

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

// 定义全局上下文中的键
type (
	ClientIPCtx struct{}
)

type Context struct {
	*gin.Context
	UserValue map[any]interface{}
}
type HandlerFunc func(c *Context)

func HandleFunc(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := new(Context)
		c.SetUserValue(ClientIPCtx{}, ctx.ClientIP())
		c.Context = ctx
		handler(c)
	}
}

func (ctx *Context) SetUserValue(key any, value interface{}) {
	m := make(map[any]interface{})
	m[key] = value
	ctx.UserValue = m
}

// GetUserValue get the value of key.
func (ctx *Context) GetUserValue(key string) interface{} {
	return ctx.UserValue[key]
}
func (ctx *Context) Success(data interface{}) {
	ret := gin.H{
		"code": 0,
		"msg":  "success",
		"time": time.Now().Unix(),
	}
	if data != nil {
		ret["data"] = &data
	}
	ctx.JSON(http.StatusOK, ret)
}
func (ctx *Context) Bind(data interface{}) error {
	err := ctx.ShouldBind(data)
	if err != nil {
		return errors.New("请求数据解析错误")
	}
	return err
}
func (ctx *Context) Error(err error) {
	statusCode := http.StatusBadRequest
	ret := gin.H{
		"code": http.StatusBadRequest,
		"msg":  err.Error(),
		"time": time.Now().Unix(),
	}
	if e, ok := err.(validation.InternalError); ok {
		ret["code"] = http.StatusInternalServerError
		ret["msg"] = "数据验证不通过"
		ret["err"] = e.Error()
		statusCode = http.StatusInternalServerError
	}
	if e, ok := err.(validation.Errors); ok {
		ret["code"] = http.StatusUnprocessableEntity
		ret["msg"] = "数据验证不通过"
		ret["err"] = e.Error()
		statusCode = http.StatusUnprocessableEntity
	}
	ctx.JSON(statusCode, ret)
}
