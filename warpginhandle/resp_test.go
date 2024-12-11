package warpginhandle_test

import "github.com/gin-gonic/gin"

// ExampleResponse 这个只是自定义的消息，在项目中你可以使用自定义的消息，这里只是提供个简单的样例
type ExampleResponse struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

func NewResponse[RES any](ctx *gin.Context, res RES, erx error) *ExampleResponse {
	return NewResp[RES](ctx, &res, erx)
}

func NewResp[RES any](ctx *gin.Context, res *RES, erx error) *ExampleResponse {
	if erx != nil {
		return &ExampleResponse{
			Code: -1,
			Desc: erx.Error(),
			Data: nil,
		}
	} else {
		return &ExampleResponse{
			Code: 0,
			Desc: "SUCCESS",
			Data: res,
		}
	}
}
