package main

import (
	"fmt"
	ginLog "gin-learn/service/user/api/internal/log"
	"gin-learn/service/user/api/internal/middleware"
	"gin-learn/service/user/api/internal/request"
	"gin-learn/service/user/api/internal/route"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.New()
	logStruct := ginLog.NewGinLog()
	logStruct.InitLog().FormatLog(server)
	middleware.GlobalMiddleware(server)
	request.Validator()
	server = route.AddRoute(server)
	err := server.Run()
	if err != nil {
		fmt.Println("服务启动失败======")
		log.Fatal(err.Error())
	}
}
