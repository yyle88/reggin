package reghan

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/erero"
)

type Handle0cFunc[RES any] func(ctx *gin.Context) (*RES, error)
type Handle1cFunc[ARG, RES any] func(ctx *gin.Context, arg *ARG) (*RES, error)

//type MakeRespFunc[RES any, RESPONSE any] func(ctx *gin.Context, res *RES, erx error) *RESPONSE //使用指针类型拼返回值

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
			//出错时就没有返回值啦
			ctx.SecureJSON(http.StatusOK, respFunc(ctx, nil, erero.WithMessage(erx, "PARAM IS WRONG")))
			return
		}
		res, erx := run(ctx, arg)
		ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erx))
	}
}

func HandleXc[ARG, RES any, RESPONSE any](run Handle1cFunc[ARG, RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle1c[ARG, RES, RESPONSE](run, BindJson[ARG], respFunc)
}
