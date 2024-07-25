package reghan

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/erero"
)

type Handle0pFunc[RES any] func() (RES, error)              //当没有参数时，比如有的接口就是没有参数的，使用这个接口注册路由
type Handle1pFunc[ARG, RES any] func(arg *ARG) (RES, error) //通常情况下请求都需要一个参数，就使用这个接口注册路由

func Handle0p[RES any, RESPONSE any](run Handle0pFunc[RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, erx := run()
		ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erx))
	}
}

func Handle1p[ARG, RES any, RESPONSE any](run Handle1pFunc[ARG, RES], parseReq ParseReqFunc[ARG], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		arg, erx := parseReq(ctx)
		if erx != nil {
			var res RES // zero
			ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erero.WithMessage(erx, "PARAM IS WRONG")))
			return
		}
		res, erx := run(arg)
		ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erx))
	}
}

func P0[RES any, RESPONSE any](run Handle0pFunc[RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle0p(run, respFunc)
}

func P1[ARG, RES any, RESPONSE any](run Handle1pFunc[ARG, RES], parseReq ParseReqFunc[ARG], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle1p(run, parseReq, respFunc)
}

func PX[ARG, RES any, RESPONSE any](run Handle1pFunc[ARG, RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle1p[ARG, RES, RESPONSE](run, BIND[ARG], respFunc)
}
