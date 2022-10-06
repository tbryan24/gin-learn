package main

import (
	clog "gin-learn/log"
	"gin-learn/middleware"
	"gin-learn/request"
	"gin-learn/route"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	clog.InitLog()
	server := gin.New()
	clog.FormatLog(server)
	middleware.GlobalMiddleware(server)
	request.Validator()
	server = route.AddRoute(server)
	err := server.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
