package context

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

// 定义全局上下文中的键
type (
	ClientIPCtx  struct{}
	transCtx     struct{}
	noTransCtx   struct{}
	transLockCtx struct{}
	userIDCtx    struct{}
	tokenCtx     struct{}
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

// NewTrans 创建事务的上下文
func NewTrans(ctx *Context, trans interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}

// FromTrans 从上下文中获取事务
func FromTrans(ctx *Context) (interface{}, bool) {
	v := ctx.Value(transCtx{})
	return v, v != nil
}

// NewNoTrans 创建不使用事务的上下文
func NewNoTrans(ctx *Context) context.Context {
	return context.WithValue(ctx, noTransCtx{}, true)
}

// FromNoTrans 从上下文中获取不使用事务标识
func FromNoTrans(ctx *Context) bool {
	v := ctx.Value(noTransCtx{})
	return v != nil && v.(bool)
}

// NewTransLock 创建事务锁的上下文
func NewTransLock(ctx *Context) context.Context {
	return context.WithValue(ctx, transLockCtx{}, true)
}

// FromTransLock 从上下文中获取事务锁
func FromTransLock(ctx *Context) bool {
	v := ctx.Value(transLockCtx{})
	return v != nil && v.(bool)
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
		"time":    time.Now(),
	}
	if data != nil {
		ret["data"] = &data
	}
	ctx.JSON(http.StatusOK, ret)
}
func (ctx *Context) Bind(data interface{}) error {
	if ctx.Request.ContentLength == 0 {
		return errors.New("The request body is empty")
	}
	return ctx.ShouldBind(data)
}
func (ctx *Context) Error(err error) {
	statusCode := http.StatusBadRequest
	ret := gin.H{
		"code":    http.StatusBadRequest,
		"message": err.Error(),
		"time":    time.Now(),
	}
	switch e := err.(type) {
	case *mysql.MySQLError:
		ret["code"] = http.StatusInternalServerError
		ret["message"] = "服务异常"
		ret["err"] = e.Error()
		statusCode = http.StatusInternalServerError
	case *json.UnmarshalTypeError:
		ret["code"] = http.StatusBadRequest
		ret["message"] = "数据验证不通过"
		ret["err"] = fmt.Sprintf("param %s should be %s not %s", e.Field, e.Type.String(), e.Value)
		statusCode = http.StatusBadRequest
	default:
		ret["code"] = http.StatusBadRequest
		ret["message"] = "数据验证不通过"
		ret["err"] = e.Error()
		statusCode = http.StatusBadRequest
	}
	ctx.JSON(statusCode, ret)
}
