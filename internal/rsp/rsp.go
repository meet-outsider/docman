package rsp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Msg  string         `json:"msg"`
	Data map[string]any `json:"data"`
}

func Build(msg string, args ...any) *Response {
	m := make(map[string]any)
	if len(args)%2 != 0 {
		args = append(args, nil)
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

func Fail(c *gin.Context, msg string, data ...any) {
	c.JSON(http.StatusInternalServerError, Build(msg, data...))
}
