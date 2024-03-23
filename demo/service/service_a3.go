package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/regginroute"
	"github.com/yyle88/regginroute/demo/message"
)

type A3 struct{}

func (a *A3) GetRoutes() regginroute.Routes[message.Response] {
	return regginroute.Routes[message.Response]{
		{Method: regginroute.GET, Path: "demo", Handle: func(c *gin.Context) message.Response {
			panic("not implemented")
		}},
		{Method: regginroute.POST, Path: "demo", Handle: func(c *gin.Context) message.Response {
			panic("not implemented")
		}},
	}
}
