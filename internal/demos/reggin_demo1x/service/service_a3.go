package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/reggin"
	"github.com/yyle88/reggin/internal/demos/reggin_demo1x/message"
)

type A3 struct{}

func (a *A3) GetRoutes() reggin.Routes[message.Response] {
	return reggin.Routes[message.Response]{
		{Method: reggin.GET, Path: "demo", Handle: func(c *gin.Context) message.Response {
			panic("not implemented")
		}},
		{Method: reggin.POST, Path: "demo", Handle: func(c *gin.Context) message.Response {
			panic("not implemented")
		}},
	}
}
