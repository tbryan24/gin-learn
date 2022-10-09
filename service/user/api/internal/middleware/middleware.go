package middleware

import (
	"github.com/gin-gonic/gin"
)

func GlobalMiddleware(server *gin.Engine) {
	//自定义中间件RequestInfos
	//server.Use(RequestInfos())
	server.Use(gin.Recovery())
}

func RequestInfosMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		//前置中间件
		//fmt.Println("请求Path: ", context.FullPath(), "请求method: ", context.Request.Method, "状态码: ", context.Writer.Status())
		context.Next()
		//后置中间件
	}
}
