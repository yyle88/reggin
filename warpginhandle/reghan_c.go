package warpginhandle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/erero"
)

// Handle0cFunc 适用于没有参数且有ctx的处理函数的场景，推荐使用
type Handle0cFunc[RES any] func(ctx *gin.Context) (RES, error)

// Handle1cFunc 适用于一个参数且有ctx的处理函数的场景，推荐使用
type Handle1cFunc[ARG, RES any] func(ctx *gin.Context, arg *ARG) (RES, error)

func Handle0c[RES any, RESPONSE any](run Handle0cFunc[RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, erx := run(ctx) //区别只在这里，这个传ctx信息，在新的开发规范中ctx还是很有用的，因此推荐使用带ctx的函数
		ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erx))
	}
}

func Handle1c[ARG, RES any, RESPONSE any](run Handle1cFunc[ARG, RES], parseReq ParseArgFunc[ARG], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		arg, erx := parseReq(ctx)
		if erx != nil {
			var res RES // zero
			ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erero.WithMessage(erx, "PARAM IS WRONG")))
			return
		}
		res, erx := run(ctx, arg) //区别只在这里，这个传ctx信息，在新的开发规范中ctx还是很有用的，因此推荐使用带ctx的函数
		ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erx))
	}
}

func C0[RES any, RESPONSE any](run Handle0cFunc[RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle0c(run, respFunc)
}

func C1[ARG, RES any, RESPONSE any](run Handle1cFunc[ARG, RES], parseReq ParseArgFunc[ARG], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle1c(run, parseReq, respFunc)
}

func CX[ARG, RES any, RESPONSE any](run Handle1cFunc[ARG, RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle1c[ARG, RES, RESPONSE](run, BIND[ARG], respFunc)
}
