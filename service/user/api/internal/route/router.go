package route

import (
	handle2 "gin-learn/service/user/api/internal/handle"
	"gin-learn/service/user/api/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRoute(server *gin.Engine) *gin.Engine {
	v1 := server.Group("/v1", middleware.RequestInfosMiddleware())
	{
		v1.Handle(http.MethodGet, "/hello", handle2.HelloHandle)
		v1.Handle(http.MethodGet, "/hello1", handle2.HelloHandle)
	}
	v2 := server.Group("/v2")
	{
		v2.GET("/test", handle2.TestHandle)
	}
	return server
}
