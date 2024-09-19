package warpginhandle

import "github.com/gin-gonic/gin"

// ResponseExample 这个只是自定义的消息，在项目中你可以使用自定义的消息，这里只是提供个简单的样例
type ResponseExample struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

func NewResponseExample[RES any](ctx *gin.Context, res RES, erx error) *ResponseExample {
	if erx != nil {
		return &ResponseExample{
			Code: -1,
			Desc: erx.Error(),
			Data: nil,
		}
	} else {
		return &ResponseExample{
			Code: 0,
			Desc: "SUCCESS",
			Data: res,
		}
	}
}

func GinResponseExample[RES any](ctx *gin.Context, res *RES, erx error) *ResponseExample {
	if erx != nil {
		return &ResponseExample{
			Code: -1,
			Desc: erx.Error(),
			Data: nil,
		}
	} else {
		return &ResponseExample{
			Code: 0,
			Desc: "SUCCESS",
			Data: res,
		}
	}
}

func newRExample[RES any](ctx *gin.Context, res RES, erx error) *ResponseExample {
	return NewResponseExample(ctx, res, erx)
}

func newPExample[RES any](ctx *gin.Context, res *RES, erx error) *ResponseExample {
	return GinResponseExample(ctx, res, erx)
}
