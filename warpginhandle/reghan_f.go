package warpginhandle

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/erero"
)

// HandleF0Func 适用于没有参数且有ctx的处理函数的场景，推荐使用
type HandleF0Func[RES any] func(ctx *gin.Context) (RES, error)

// HandleF1Func 适用于一个参数且有ctx的处理函数的场景，推荐使用
type HandleF1Func[ARG, RES any] func(ctx *gin.Context, arg *ARG) (RES, error)

func HandleF0[RES any, RESPONSE any](run HandleF0Func[RES], errorMsg ErrorMsgFunc[RESPONSE], status *StatusConfig) gin.HandlerFunc {
	status.MustNice()
	return func(ctx *gin.Context) {
		res, err := run(ctx) //区别只在这里，这个不传ctx信息，因此处理逻辑里拿不到ctx信息，适用于简单场景
		if err != nil {
			ctx.SecureJSON(status.BadLogic, errorMsg(ctx, err))
			return
		}
		ctx.SecureJSON(status.StatusOK, res)
	}
}

func HandleF1[ARG, RES any, RESPONSE any](run HandleF1Func[ARG, RES], parseReq ParseArgFunc[ARG], errorMsg ErrorMsgFunc[RESPONSE], status *StatusConfig) gin.HandlerFunc {
	status.MustNice()
	return func(ctx *gin.Context) {
		arg, err := parseReq(ctx)
		if err != nil {
			// var res RES // zero
			ctx.SecureJSON(status.BadParam, errorMsg(ctx, erero.WithMessage(err, "PARAM IS WRONG")))
			return
		}
		res, err := run(ctx, arg) //区别只在这里，这个不传ctx信息，因此处理逻辑里拿不到ctx信息，适用于简单场景
		if err != nil {
			ctx.SecureJSON(status.BadLogic, errorMsg(ctx, err))
			return
		}
		ctx.SecureJSON(status.StatusOK, res)
	}
}

func F0[RES any, RESPONSE any](run HandleF0Func[RES], errorMsg ErrorMsgFunc[RESPONSE]) gin.HandlerFunc {
	return HandleF0(run, errorMsg, NewStatus400())
}

func F1[ARG, RES any, RESPONSE any](run HandleF1Func[ARG, RES], parseReq ParseArgFunc[ARG], errorMsg ErrorMsgFunc[RESPONSE]) gin.HandlerFunc {
	return HandleF1(run, parseReq, errorMsg, NewStatus400())
}

func FX[ARG, RES any, RESPONSE any](run HandleF1Func[ARG, RES], errorMsg ErrorMsgFunc[RESPONSE]) gin.HandlerFunc {
	return HandleF1[ARG, RES, RESPONSE](run, BIND[ARG], errorMsg, NewStatus400())
}
