package log

import (
	"fmt"
	"gin-learn/common/utils"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

type GinLog struct{}

func NewGinLog() *GinLog {
	return &GinLog{}
}

func (*GinLog) InitLog() *GinLog {
	// 禁用控制台颜色
	//gin.DisableConsoleColor()
	dirPath := "./logs/user/api/"
	fileName := "gin.log"
	res, _ := utils.PathExists(dirPath)
	if !res {
		utils.CreateMultiDir(dirPath)
	}
	// 创建记录日志的文件
	f, _ := os.Create(dirPath + fileName)
	//gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	return &GinLog{}
}

func (*GinLog) FormatLog(server *gin.Engine) *GinLog {
	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	return &GinLog{}
}
