package warpginhandle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/must"
)

// MakeRespFunc 这个函数能把结果再包装成带附加信息的，比如返回带错误码或者再从ctx里取东西，做拼装/翻译等操作
type MakeRespFunc[RES any, RESPONSE any] func(ctx *gin.Context, res RES, cause error) *RESPONSE

// ErrorMsgFunc 这个函数把错误提示转换为带附加信息的
type ErrorMsgFunc[RESPONSE any] func(ctx *gin.Context, cause error) *RESPONSE

type StatusConfig struct {
	StatusOK int
	BadParam int
	BadLogic int
}

func NewStatus200() *StatusConfig {
	return &StatusConfig{
		StatusOK: http.StatusOK,
		BadParam: http.StatusOK,
		BadLogic: http.StatusOK,
	}
}

func NewStatus400() *StatusConfig {
	return &StatusConfig{
		StatusOK: http.StatusOK,
		BadParam: http.StatusBadRequest,
		BadLogic: http.StatusBadRequest,
	}
}

func (status *StatusConfig) MustNice() {
	must.Full(status)
	must.Nice(status.StatusOK)
	must.Nice(status.BadParam)
	must.Nice(status.BadLogic)
}
