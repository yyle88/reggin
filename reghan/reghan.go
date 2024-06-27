package reghan

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yyle88/erero"
)

type Handle0pFunc[RES any] func() (*RES, error)
type Handle1pFunc[ARG, RES any] func(arg *ARG) (*RES, error)

type MakeRespFunc[RES any, RESPONSE any] func(res *RES, erx error) *RESPONSE
type ParseReqFunc[ARG any] func(c *gin.Context) (*ARG, error)

func Handle0p[RES any, RESPONSE any](run Handle0pFunc[RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.SecureJSON(http.StatusOK, respFunc(run()))
	}
}

func Handle1p[ARG, RES any, RESPONSE any](run Handle1pFunc[ARG, RES], parseReq ParseReqFunc[ARG], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		arg, erx := parseReq(ctx)
		if erx != nil {
			ctx.SecureJSON(http.StatusOK, respFunc(nil, erero.WithMessage(erx, "PARAM IS WRONG")))
			return
		}
		ctx.SecureJSON(http.StatusOK, respFunc(run(arg)))
	}
}

func Handle1x[ARG, RES any, RESPONSE any](run Handle1pFunc[ARG, RES], respFunc MakeRespFunc[RES, RESPONSE]) gin.HandlerFunc {
	return Handle1p[ARG, RES, RESPONSE](run, BindJson[ARG], respFunc)
}

func BindJson[ARG any](ctx *gin.Context) (arg *ARG, err error) {
	var req ARG
	if erx := ctx.ShouldBindBodyWith(&req, binding.JSON); erx != nil {
		return nil, erero.WithMessage(erx, "CAN NOT BIND REQ")
	}
	return &req, nil
}
