package reghan

import "github.com/gin-gonic/gin"

// ResponseType 这个只是自定义的消息，当然实际上它不和任何逻辑挂钩
// 在项目中你可以使用自定义的消息，这里只是提供个简单的样例
type ResponseType struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

func NewResponse[resType any](ctx *gin.Context, res resType, erx error) *ResponseType {
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

func MakeResponse[resType any](ctx *gin.Context, res *resType, erx error) *ResponseType {
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
