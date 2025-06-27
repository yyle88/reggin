package warpginhandle_test

import "github.com/gin-gonic/gin"

// 这个只是自定义的消息，在项目中你可以使用自定义的消息，这里只是提供个简单的样例
type respType struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

func makeResp[RES any](ctx *gin.Context, res RES, cause error) *respType {
	return warpResp[RES](ctx, &res, cause)
}

func warpResp[RES any](ctx *gin.Context, res *RES, cause error) *respType {
	if cause != nil {
		return wrongResp(ctx, cause)
	} else {
		return rightResp(ctx, res)
	}
}

func rightResp[RES any](ctx *gin.Context, res *RES) *respType {
	return &respType{
		Code: 0,
		Desc: "SUCCESS",
		Data: res,
	}
}

func wrongResp(ctx *gin.Context, cause error) *respType {
	return &respType{
		Code: -1,
		Desc: cause.Error(),
		Data: nil,
	}
}
