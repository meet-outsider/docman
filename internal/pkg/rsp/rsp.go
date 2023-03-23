package rsp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Build(msg string, args ...any) *Response {
	if args == nil {
		return &Response{
			Msg: msg,
		}
	}
	m := make(map[string]any)
	if len(args) == 1 {
		return &Response{
			Msg:  msg,
			Data: args,
		}
	}
	if len(args)%2 != 0 {
		return &Response{
			Msg:  "build error,please check your args quantity",
			Data: nil,
		}
	}
	for i := 0; i < len(args); i += 2 {
		if key, ok := args[i].(string); ok {
			m[key] = args[i+1]
		}
	}
	return &Response{
		Msg:  msg,
		Data: m,
	}
}

func Ok(c *gin.Context, msg string, data ...any) {
	c.JSON(http.StatusOK, Build(msg, data...))
}
func Created(c *gin.Context, msg string, data ...any) {
	c.JSON(http.StatusCreated, Build(msg, data...))
}

func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, Build(msg))
}
func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, Build(msg))
}
func Custom(c *gin.Context, status int, msg string, data ...any) {
	c.JSON(status, Build(msg, data...))
}
