package server

import (
	"github.com/gin-gonic/gin"
)

var Inst *gin.Engine

func GET(relativePath string, handlers func(*gin.Context)) gin.IRoutes {
	return Inst.GET(relativePath, handlers)
}

func POST(relativePath string, handlers func(*gin.Context)) gin.IRoutes {
	return Inst.POST(relativePath, handlers)
}
func PUT(relativePath string, handlers func(*gin.Context)) gin.IRoutes {
	return Inst.PUT(relativePath, handlers)
}
func DELETE(relativePath string, handlers func(*gin.Context)) gin.IRoutes {
	return Inst.DELETE(relativePath, handlers)
}
func PATCH(relativePath string, handlers func(*gin.Context)) gin.IRoutes {
	return Inst.PATCH(relativePath, handlers)
}
