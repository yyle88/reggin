package reghan

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/erero"
)

type Handle0aFunc[RES any] func() (RES, error)              //当需要返回非指针类型时，比如 int/string/bool/float64 这些基本类型
type Handle1aFunc[ARG, RES any] func(arg *ARG) (RES, error) //使用基本类型做返回值，这时候结果也最好是基本类型，而非指针类型

type MakeRespBase[RES any, RESPONSE any] func(ctx *gin.Context, res RES, erx error) *RESPONSE //使用基本类型做返回值

func Handle0a[RES any, RESPONSE any](run Handle0aFunc[RES], respFunc MakeRespBase[RES, RESPONSE]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, erx := run()
		ctx.SecureJSON(http.StatusOK, respFunc(ctx, res, erx))
	}
}

func Handle1a[ARG, RES any, RESPONSE any](run Handle1aFunc[ARG, RES], parseReq ParseReqFunc[ARG], respFunc MakeRespBase[RES, RESPONSE]) gin.HandlerFunc {
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

func HandleXa[ARG, RES any, RESPONSE any](run Handle1aFunc[ARG, RES], respFunc MakeRespBase[RES, RESPONSE]) gin.HandlerFunc {
	return Handle1a[ARG, RES, RESPONSE](run, BindJson[ARG], respFunc)
}
