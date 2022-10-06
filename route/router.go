package route

import (
	"gin-learn/handle"
	"gin-learn/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRoute(server *gin.Engine) *gin.Engine {
	v1 := server.Group("/v1", middleware.RequestInfosMiddleware())
	{
		v1.Handle(http.MethodGet, "/hello", handle.HelloHandle)
		v1.Handle(http.MethodGet, "/hello1", handle.HelloHandle)
	}
	v2 := server.Group("/v2")
	{
		v2.GET("/test", handle.TestHandle)
	}
	return server
}
