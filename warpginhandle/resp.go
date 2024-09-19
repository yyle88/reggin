package warpginhandle

import "github.com/gin-gonic/gin"

// MakeRespFunc 这个函数能把结果再包装成带附加信息的，比如返回带错误码或者再从ctx里取东西，做拼装/翻译等操作
type MakeRespFunc[RES any, RESPONSE any] func(ctx *gin.Context, res RES, erx error) *RESPONSE
