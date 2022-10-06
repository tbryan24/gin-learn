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
		//path := context.FullPath()
		//method := context.Request.Method
		//fmt.Println("请求Path: ", path)
		//fmt.Println("请求method: ", method)
		context.Next()
		//fmt.Println("状态码: ", context.Writer.Status())
	}
}
