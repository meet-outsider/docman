package server

import (
	"github.com/gin-gonic/gin"
)

var G *gin.Engine

func GET(relativePath string, handlers func(*gin.Context)) gin.IRoutes {
	return G.GET(relativePath, handlers)
}

func POST(relativePath string, handlers func(*gin.Context)) gin.IRoutes {
	return G.POST(relativePath, handlers)
}
func PUT(relativePath string, handlers func(*gin.Context)) gin.IRoutes {
	return G.PUT(relativePath, handlers)
}
func DELETE(relativePath string, handlers func(*gin.Context)) gin.IRoutes {
	return G.DELETE(relativePath, handlers)
}
func PATCH(relativePath string, handlers func(*gin.Context)) gin.IRoutes {
	return G.PATCH(relativePath, handlers)
}
