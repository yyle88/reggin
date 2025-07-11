package warpginhandle

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/erero"
)

// HandleP0Func 适用于没有参数且无ctx的处理函数的场景，认为不带ctx的属于非正式的场景，没法拿到上下文的信息，比如监控或者超时等信息，但比较简单
type HandleP0Func[RES any] func() (RES, error)

// HandleP1Func 适用于一个参数且无ctx的处理函数的场景
type HandleP1Func[ARG, RES any] func(arg *ARG) (RES, error)

func HandleP0[RES any, RESPONSE any](run HandleP0Func[RES], makeResp MakeRespFunc[RES, RESPONSE], status *StatusConfig) gin.HandlerFunc {
	status.MustNice()
	return func(ctx *gin.Context) {
		res, err := run() //区别只在这里，这个不传ctx信息，因此处理逻辑里拿不到ctx信息，适用于简单场景
		if err != nil {
			ctx.SecureJSON(status.BadLogic, makeResp(ctx, res, err))
			return
		}
		ctx.SecureJSON(status.StatusOK, makeResp(ctx, res, nil))
	}
}

func HandleP1[ARG, RES any, RESPONSE any](run HandleP1Func[ARG, RES], parseReq ParseArgFunc[ARG], makeResp MakeRespFunc[RES, RESPONSE], status *StatusConfig) gin.HandlerFunc {
	status.MustNice()
	return func(ctx *gin.Context) {
		arg, err := parseReq(ctx)
		if err != nil {
			var res RES // zero
			ctx.SecureJSON(status.BadParam, makeResp(ctx, res, erero.WithMessage(err, "PARAM IS WRONG")))
			return
		}
		res, err := run(arg) //区别只在这里，这个不传ctx信息，因此处理逻辑里拿不到ctx信息，适用于简单场景
		if err != nil {
			ctx.SecureJSON(status.BadLogic, makeResp(ctx, res, err))
			return
		}
		ctx.SecureJSON(status.StatusOK, makeResp(ctx, res, nil))
	}
}

func P0[RES any, RESPONSE any](run HandleP0Func[RES], makeResp MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return HandleP0(run, makeResp, NewStatus200())
}

func P1[ARG, RES any, RESPONSE any](run HandleP1Func[ARG, RES], parseReq ParseArgFunc[ARG], makeResp MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return HandleP1(run, parseReq, makeResp, NewStatus200())
}

func PX[ARG, RES any, RESPONSE any](run HandleP1Func[ARG, RES], makeResp MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return HandleP1[ARG, RES, RESPONSE](run, BIND[ARG], makeResp, NewStatus200())
}
