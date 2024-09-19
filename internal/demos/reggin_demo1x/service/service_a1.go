package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/reggin"
	"github.com/yyle88/reggin/internal/demos/reggin_demo1x/message"
)

// A1 define a service
// 自定义的服务
type A1 struct{}

// GetRoutes defile routers
// 自定义的接口
func (a *A1) GetRoutes() reggin.Routes[message.Response] {
	return reggin.Routes[message.Response]{
		{Method: reggin.GET, Path: "demo", Handle: a.HandleGetDemo},
		{Method: reggin.POST, Path: "demo", Handle: a.HandlePostDemo},
		{Method: reggin.POST, Path: "set", Handle: a.HandleSetDemo},
	}
}

func (a *A1) HandleGetDemo(c *gin.Context) message.Response {
	//write logic here 在这里实现你的逻辑
	return message.Response{
		Code: 0,
		Desc: "OK",
		Data: map[string]any{"a": "a", "b": "b", "c": "c"},
	}
}

func (a *A1) HandlePostDemo(c *gin.Context) message.Response {
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

func (a *A1) HandleSetDemo(c *gin.Context) message.Response {
	type requestType struct {
		X int
	}
	type responseType struct {
		Y int
	}
	var req requestType
	if err := c.ShouldBindJSON(&req); err != nil {
		return message.Response{
			Code: -1,
			Desc: "wrong param",
			Data: nil,
		}
	}
	//write some logic
	res := req.X * 2
	//set return value
	return message.Response{
		Code: 0,
		Desc: "OK",
		Data: responseType{Y: res},
	}
}
