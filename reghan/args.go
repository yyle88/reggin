package reghan

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yyle88/erero"
)

type ParseReqFunc[ARG any] func(c *gin.Context) (*ARG, error)

func BindJson[ARG any](ctx *gin.Context) (arg *ARG, err error) {
	var req ARG
	if erx := ctx.ShouldBindBodyWith(&req, binding.JSON); erx != nil {
		return nil, erero.WithMessage(erx, "CAN NOT BIND REQ")
	}
	return &req, nil
}
