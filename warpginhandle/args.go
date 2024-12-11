package warpginhandle

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/schema"
	"github.com/yyle88/erero"
)

type ParseArgFunc[ARG any] func(c *gin.Context) (*ARG, error)

var (
	schemaFormDecoder = schema.NewDecoder()
	schemaJsonDecoder = schema.NewDecoder()
)

func init() {
	schemaFormDecoder.SetAliasTag("form")
	schemaJsonDecoder.SetAliasTag("json")
}

// BIND 绑定参数，这个函数名太长其实不太有利于使用，但有时为了逻辑清晰也可以用它
func BIND[ARG any](ctx *gin.Context) (arg *ARG, err error) {
	var req ARG
	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		return nil, erero.WithMessage(err, "CAN NOT BIND REQ")
	}
	return &req, nil
}

// B 也是绑定参数，在使用时这块的字符越少越有利于突出泛型本身的类型，因此使用这个很短的函数名
func B[ARG any](ctx *gin.Context) (arg *ARG, err error) {
	var req ARG
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, erero.WithMessage(err, "CAN NOT BIND REQ")
	}
	return &req, nil
}

// Q 也是绑定参数，只是绑定的是 uri param，也就是 GET 请求的参数，这时候需要定义 form 标签
func Q[ARG any](ctx *gin.Context) (arg *ARG, err error) {
	var req ARG
	if err := ctx.ShouldBindQuery(&req); err != nil {
		return nil, erero.WithMessage(err, "CAN NOT BIND REQ")
	}
	return &req, nil
}

// QueryForm 这个和前面 Q 函数的功能是一样的，因为 gin 框架默认的 QueryParams 就是用 form 解析的，因此这个不建议用，就用 Q 就行
func QueryForm[ARG any](ctx *gin.Context) (arg *ARG, err error) {
	var req ARG
	if err := schemaFormDecoder.Decode(&req, ctx.Request.URL.Query()); err != nil {
		return nil, erero.WithMessage(err, "CAN NOT DECODE FORM FROM REQ")
	}
	return &req, nil
}

// QueryJson 假如是 GET 请求，但你依然想用 json 标签去收，你就可以用这个函数，让你能用 json 标签收 GET 的 QueryParams 请求参数
func QueryJson[ARG any](ctx *gin.Context) (arg *ARG, err error) {
	var req ARG
	if err := schemaJsonDecoder.Decode(&req, ctx.Request.URL.Query()); err != nil {
		return nil, erero.WithMessage(err, "CAN NOT DECODE JSON FROM REQ")
	}
	return &req, nil
}
