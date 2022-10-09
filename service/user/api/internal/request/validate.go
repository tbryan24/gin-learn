package request

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Validator() {
	// 3、将我们自定义的校验方法注册到 validator中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("NotNullAndAdmin", nameNotNullAndAdmin)
		if err != nil {
			panic("RegisterValidation error")
		}
	}
}
