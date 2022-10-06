package request

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type HelloRequest struct {
	Name string `json:"name" binding:"required,NotNullAndAdmin"`
	Age  int64  `json:"age" binding:"required,gt=10"`
}

// 自定义验证规则断言
var nameNotNullAndAdmin validator.Func = func(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		fmt.Println(value, 11)
		// 字段不能为空，并且不等于  admin
		return value != "" && !("admin" == value)
	}
	return true
}

type HelloResponse struct {
	Status  int64
	Message string
	Data    interface{}
}
