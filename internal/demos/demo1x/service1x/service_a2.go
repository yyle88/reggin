package service1x

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/reggin"
	"github.com/yyle88/reggin/internal/demos/demo1x/message1x"
)

type A2 struct{}

func (a *A2) GetRoutes() reggin.Routes[message1x.Response] {
	return reggin.Routes[message1x.Response]{
		{Method: reggin.GET, Path: "demo", Handle: a.HandleGetDemo},
		{Method: reggin.POST, Path: "demo", Handle: a.HandlePostDemo},
	}
}

func (a *A2) HandleGetDemo(c *gin.Context) message1x.Response {
	panic("not implemented")
}

func (a *A2) HandlePostDemo(c *gin.Context) message1x.Response {
	panic("not implemented")
}
