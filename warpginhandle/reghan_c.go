package warpginhandle

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/erero"
)

// HandleC0Func 适用于没有参数且有ctx的处理函数的场景，推荐使用
type HandleC0Func[RES any] func(ctx *gin.Context) (RES, error)

// HandleC1Func 适用于一个参数且有ctx的处理函数的场景，推荐使用
type HandleC1Func[ARG, RES any] func(ctx *gin.Context, arg *ARG) (RES, error)

func HandleC0[RES any, RESPONSE any](run HandleC0Func[RES], makeResp MakeRespFunc[RES, RESPONSE], status *StatusConfig) gin.HandlerFunc {
	status.MustNice()
	return func(ctx *gin.Context) {
		res, err := run(ctx) //区别只在这里，这个传ctx信息，在新的开发规范中ctx还是很有用的，因此推荐使用带ctx的函数
		if err != nil {
			ctx.SecureJSON(status.BadLogic, makeResp(ctx, res, err))
			return
		}
		ctx.SecureJSON(status.StatusOK, makeResp(ctx, res, nil))
	}
}

func HandleC1[ARG, RES any, RESPONSE any](run HandleC1Func[ARG, RES], parseReq ParseArgFunc[ARG], makeResp MakeRespFunc[RES, RESPONSE], status *StatusConfig) gin.HandlerFunc {
	status.MustNice()
	return func(ctx *gin.Context) {
		arg, err := parseReq(ctx)
		if err != nil {
			var res RES // zero
			ctx.SecureJSON(status.BadParam, makeResp(ctx, res, erero.WithMessage(err, "PARAM IS WRONG")))
			return
		}
		res, err := run(ctx, arg) //区别只在这里，这个传ctx信息，在新的开发规范中ctx还是很有用的，因此推荐使用带ctx的函数
		if err != nil {
			ctx.SecureJSON(status.BadLogic, makeResp(ctx, res, err))
			return
		}
		ctx.SecureJSON(status.StatusOK, makeResp(ctx, res, err))
	}
}

func C0[RES any, RESPONSE any](run HandleC0Func[RES], makeResp MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return HandleC0(run, makeResp, NewStatus200())
}

func C1[ARG, RES any, RESPONSE any](run HandleC1Func[ARG, RES], parseReq ParseArgFunc[ARG], makeResp MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return HandleC1(run, parseReq, makeResp, NewStatus200())
}

func CX[ARG, RES any, RESPONSE any](run HandleC1Func[ARG, RES], makeResp MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return HandleC1[ARG, RES, RESPONSE](run, BIND[ARG], makeResp, NewStatus200())
}
