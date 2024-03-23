package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/regginroute"
	"github.com/yyle88/regginroute/demo/message"
)

// A define a service
// 自定义的服务
type A struct{}

// GetRoutes defile routers
// 自定义的接口
func (a *A) GetRoutes() regginroute.Routes[message.Response] {
	return regginroute.Routes[message.Response]{
		{Method: regginroute.GET, Path: "demo", Handle: a.HandleGetDemo},
		{Method: regginroute.POST, Path: "demo", Handle: a.HandlePostDemo},
	}
}

func (a *A) HandleGetDemo(c *gin.Context) message.Response {
	//write logic here 在这里实现你的逻辑
	return message.Response{
		Code: 0,
		Desc: "OK",
		Data: map[string]any{"a": "a", "b": "b", "c": "c"},
	}
}

func (a *A) HandlePostDemo(c *gin.Context) message.Response {
	//write logic here 在这里实现你的逻辑
	//example:
	//defile request type. define it here is more clean than defile it outside the func.
	type requestType struct {
		X int
	}
	//define response data type. define it at the top of the func body can make the code very clean.
	type responseDataType struct {
		X int
		Y int
		Z int
	}
	//就是说，最好是在函数的起始阶段就定义req和resp的数据类型，接下来是解析req的数据
	var req requestType
	if err := c.ShouldBindJSON(&req); err != nil {
		return message.Response{
			Code: -1,
			Desc: "wrong param",
			Data: nil,
		}
	}
	// write logic here to get the data.
	data := responseDataType{
		X: req.X * 1,
		Y: req.X * 2,
		Z: req.X * 3,
	}
	// return the response
	return message.Response{
		Code: 0,
		Desc: "OK",
		Data: data, // set response data here.
	}
}
