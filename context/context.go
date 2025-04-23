package context

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	errors "github.com/copkg/gopkg/errors"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-playground/validator/v10"
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
		"code":    http.StatusOK,
		"message": "success",
		"time":    time.Now().Unix(),
	}
	if data != nil {
		ret["data"] = &data
	}
	ctx.JSON(http.StatusOK, ret)
}
func (ctx *Context) Bind(data interface{}) error {
	return ctx.ShouldBind(data)
}
func (ctx *Context) Error(err error) {
	statusCode := http.StatusBadRequest
	ret := gin.H{
		"code":    http.StatusBadRequest,
		"message": err.Error(),
		"time":    time.Now().Unix(),
	}
	if e, ok := err.(validation.InternalError); ok {
		ret["code"] = http.StatusInternalServerError
		ret["message"] = "数据验证不通过"
		ret["err"] = e.Error()
		statusCode = http.StatusInternalServerError
	}
	if e, ok := err.(validation.Errors); ok {
		ret["code"] = http.StatusUnprocessableEntity
		ret["message"] = "数据验证不通过"
		ret["err"] = e.Error()
		statusCode = http.StatusUnprocessableEntity
	}
	if _, ok := err.(validator.ValidationErrors); ok {
		var errmsg []string
		for _, ferr := range err.(validator.ValidationErrors) {
			errmsg = append(errmsg, fmt.Sprintf("param %s %s", strings.ToLower(ferr.Field()), ferr.Tag()))
		}
		// 将字符串数组包装在一个map中以输出JSON格式
		ret["code"] = http.StatusUnprocessableEntity
		ret["message"] = "数据验证不通过"
		ret["err"] = errmsg
		statusCode = http.StatusUnprocessableEntity
	}
	if e, ok := err.(errors.Error); ok {
		ret["code"] = e.HttpCode()
		ret["message"] = err.Error()
		ret["err"] = e.Wrap()
		statusCode = http.StatusUnprocessableEntity
	}
	ctx.JSON(statusCode, ret)
}
