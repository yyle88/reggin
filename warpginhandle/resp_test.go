package warpginhandle_test

import "github.com/gin-gonic/gin"

// ExampleResponse 这个只是自定义的消息，在项目中你可以使用自定义的消息，这里只是提供个简单的样例
type ExampleResponse struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

func NewResponse[RES any](ctx *gin.Context, res RES, err error) *ExampleResponse {
	return NewResp[RES](ctx, &res, err)
}

func NewResp[RES any](ctx *gin.Context, res *RES, err error) *ExampleResponse {
	if err != nil {
		return &ExampleResponse{
			Code: -1,
			Desc: err.Error(),
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
