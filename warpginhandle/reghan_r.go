package warpginhandle

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/erero"
)

// HandleR0Func 适用于没有参数且无ctx的处理函数的场景，认为不带ctx的属于非正式的场景，没法拿到上下文的信息，比如监控或者超时等信息，但比较简单
type HandleR0Func[RES any] func() (RES, error)

// HandleR1Func 适用于一个参数且无ctx的处理函数的场景
type HandleR1Func[ARG, RES any] func(arg *ARG) (RES, error)

func HandleR0[RES any, RESPONSE any](run HandleR0Func[RES], errorMsg ErrorMsgFunc[RESPONSE], status *StatusConfig) gin.HandlerFunc {
	status.MustNice()
	return func(ctx *gin.Context) {
		res, err := run() //区别只在这里，这个不传ctx信息，因此处理逻辑里拿不到ctx信息，适用于简单场景
		if err != nil {
			ctx.SecureJSON(status.BadLogic, errorMsg(ctx, err))
			return
		}
		ctx.SecureJSON(status.StatusOK, res)
	}
}

func HandleR1[ARG, RES any, RESPONSE any](run HandleR1Func[ARG, RES], parseReq ParseArgFunc[ARG], errorMsg ErrorMsgFunc[RESPONSE], status *StatusConfig) gin.HandlerFunc {
	status.MustNice()
	return func(ctx *gin.Context) {
		arg, err := parseReq(ctx)
		if err != nil {
			// var res RES // zero
			ctx.SecureJSON(status.BadParam, errorMsg(ctx, erero.WithMessage(err, "PARAM IS WRONG")))
			return
		}
		res, err := run(arg) //区别只在这里，这个不传ctx信息，因此处理逻辑里拿不到ctx信息，适用于简单场景
		if err != nil {
			ctx.SecureJSON(status.BadLogic, errorMsg(ctx, err))
			return
		}
		ctx.SecureJSON(status.StatusOK, res)
	}
}

func R0[RES any, RESPONSE any](run HandleR0Func[RES], errorMsg ErrorMsgFunc[RESPONSE]) gin.HandlerFunc {
	return HandleR0(run, errorMsg, NewStatus400())
}

func R1[ARG, RES any, RESPONSE any](run HandleR1Func[ARG, RES], parseReq ParseArgFunc[ARG], errorMsg ErrorMsgFunc[RESPONSE]) gin.HandlerFunc {
	return HandleR1(run, parseReq, errorMsg, NewStatus400())
}

func RX[ARG, RES any, RESPONSE any](run HandleR1Func[ARG, RES], errorMsg ErrorMsgFunc[RESPONSE]) gin.HandlerFunc {
	return HandleR1[ARG, RES, RESPONSE](run, BIND[ARG], errorMsg, NewStatus400())
}
