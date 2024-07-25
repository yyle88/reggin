package reghan

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yyle88/erero"
)

type ParseReqFunc[ARG any] func(c *gin.Context) (*ARG, error)

// BIND 绑定参数，这个函数名太长其实不太有利于使用，但有时为了逻辑清晰也可以用它
func BIND[ARG any](ctx *gin.Context) (arg *ARG, err error) {
	var req ARG
	if erx := ctx.ShouldBindBodyWith(&req, binding.JSON); erx != nil {
		return nil, erero.WithMessage(erx, "CAN NOT BIND REQ")
	}
	return &req, nil
}

// B 也是绑定参数，在使用时这块的字符越少越有利于突出泛型本身的类型，因此使用这个很短的函数名
func B[ARG any](ctx *gin.Context) (arg *ARG, err error) {
	var req ARG
	if erx := ctx.ShouldBindJSON(&req); erx != nil {
		return nil, erero.WithMessage(erx, "CAN NOT BIND REQ")
	}
	return &req, nil
}

// Q 也是绑定参数，只是绑定的是 uri param，也就是 GET 请求的参数，这时候需要定义 form 标签
func Q[ARG any](ctx *gin.Context) (arg *ARG, err error) {
	var req ARG
	if erx := ctx.ShouldBindQuery(&req); erx != nil {
		return nil, erero.WithMessage(erx, "CAN NOT BIND REQ")
	}
	return &req, nil
}
