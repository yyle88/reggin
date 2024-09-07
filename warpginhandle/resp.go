package warpginhandle

import "github.com/gin-gonic/gin"

// MakeRespFunc 这个函数能把结果再包装成带附加信息的，比如返回带错误码或者再从ctx里取东西，做拼装/翻译等操作
type MakeRespFunc[RES any, RESPONSE any] func(ctx *gin.Context, res RES, erx error) *RESPONSE

// ResponseType 这个只是自定义的消息，在项目中你可以使用自定义的消息，这里只是提供个简单的样例
type ResponseType struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

func NewResponse[RES any](ctx *gin.Context, res RES, erx error) *ResponseType {
	if erx != nil {
		return &ResponseType{
			Code: -1,
			Desc: erx.Error(),
			Data: nil,
		}
	} else {
		return &ResponseType{
			Code: 0,
			Desc: "SUCCESS",
			Data: res,
		}
	}
}

func GinResponse[RES any](ctx *gin.Context, res *RES, erx error) *ResponseType {
	if erx != nil {
		return &ResponseType{
			Code: -1,
			Desc: erx.Error(),
			Data: nil,
		}
	} else {
		return &ResponseType{
			Code: 0,
			Desc: "SUCCESS",
			Data: res,
		}
	}
}

func Rt[RES any](ctx *gin.Context, res RES, erx error) *ResponseType {
	return NewResponse(ctx, res, erx)
}

func Rp[RES any](ctx *gin.Context, res *RES, erx error) *ResponseType {
	return GinResponse(ctx, res, erx)
}
