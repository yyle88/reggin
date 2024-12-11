package service1x

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/reggin"
	"github.com/yyle88/reggin/internal/demos/demo1x/message1x"
)

type A3 struct{}

func (a *A3) GetRoutes() reggin.Routes[message1x.Response] {
	return reggin.Routes[message1x.Response]{
		{Method: reggin.GET, Path: "demo", Handle: func(c *gin.Context) message1x.Response {
			panic("not implemented")
		}},
		{Method: reggin.POST, Path: "demo", Handle: func(c *gin.Context) message1x.Response {
			panic("not implemented")
		}},
	}
}
