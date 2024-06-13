package context

import (
	"github.com/copkg/gopkg/schema"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"time"
)

type Context struct {
	*gin.Context
	UserValue map[string]interface{}
}
type HandlerFunc func(c *Context)

func HandleFunc(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := new(Context)
		c.Context = ctx
		handler(c)
	}
}

func (ctx *Context) SetUserValue(key string, value interface{}) {
	ctx.UserValue[key] = value
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
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				obj := reflect.TypeOf(data)
				if obj.Kind() == reflect.Ptr {
					obj = obj.Elem()
				}
				if f, exist := obj.FieldByName(e.Field()); exist {
					return schema.Error{
						Code: 422,
						Msg:  f.Tag.Get("error"),
						Err:  e,
					}
				}
				return schema.Error{
					Code: 422,
					Msg:  "数据验证不通过",
					Err:  e,
				}
			}
		}
		return err
	}
	return nil
}
func (ctx *Context) Error(err error) {
	statusCode := http.StatusBadRequest
	ret := gin.H{
		"code": http.StatusBadRequest,
		"msg":  err.Error(),
		"time": time.Now().Unix(),
	}
	if e, ok := err.(*schema.Error); ok {
		ret["code"] = e.Code
		ret["msg"] = e.Msg
		ret["err"] = e.Err.Error()
		statusCode = e.StatusCode
	}
	ctx.JSON(statusCode, ret)
}
