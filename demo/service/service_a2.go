package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/regginroute"
	"github.com/yyle88/regginroute/demo/message"
)

type A2 struct{}

func (a *A2) GetRoutes() regginroute.Routes[message.Response] {
	return regginroute.Routes[message.Response]{
		{Method: regginroute.GET, Path: "demo", Handle: a.HandleGetDemo},
		{Method: regginroute.POST, Path: "demo", Handle: a.HandlePostDemo},
	}
}

func (a *A2) HandleGetDemo(c *gin.Context) message.Response {
	panic("not implemented")
}

func (a *A2) HandlePostDemo(c *gin.Context) message.Response {
	panic("not implemented")
}
