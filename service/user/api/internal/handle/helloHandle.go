package handle

import (
	"gin-learn/service/user/api/internal/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandle(context *gin.Context) {
	var HelloRequest request.HelloRequest
	if err := context.ShouldBindJSON(&HelloRequest); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	resp := request.HelloResponse{Status: 1, Message: "success", Data: gin.H{
		"name": HelloRequest.Name,
		"age":  HelloRequest.Age,
	}}
	context.JSON(http.StatusOK, &resp)
}
