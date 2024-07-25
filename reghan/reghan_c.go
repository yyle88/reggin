package reghan

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/erero"
)

type Handle0cFunc[RES any] func(ctx *gin.Context) (RES, error)                //当需要返回非指针类型时，比如 int/string/bool/float64 这些基本类型
type Handle1cFunc[ARG, RES any] func(ctx *gin.Context, arg *ARG) (RES, error) //使用基本类型做返回值，这时候结果也最好是基本类型，而非指针类型

func Handle0c[RES any, RESPONSE any](run Handle0cFunc[RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, erx := run(ctx)
		ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erx))
	}
}

func Handle1c[ARG, RES any, RESPONSE any](run Handle1cFunc[ARG, RES], parseReq ParseReqFunc[ARG], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		arg, erx := parseReq(ctx)
		if erx != nil {
			var res RES // zero
			ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erero.WithMessage(erx, "PARAM IS WRONG")))
			return
		}
		res, erx := run(ctx, arg)
		ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erx))
	}
}

func C0[RES any, RESPONSE any](run Handle0cFunc[RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle0c(run, respFunc)
}

func C1[ARG, RES any, RESPONSE any](run Handle1cFunc[ARG, RES], parseReq ParseReqFunc[ARG], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle1c(run, parseReq, respFunc)
}

func CX[ARG, RES any, RESPONSE any](run Handle1cFunc[ARG, RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle1c[ARG, RES, RESPONSE](run, BIND[ARG], respFunc)
}
