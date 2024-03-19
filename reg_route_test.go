package regginroute

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	resty2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/regginroute/utilsregginroute"
)

func TestMain(m *testing.M) {
	g := gin.New()
	g.Use(gin.Recovery())
	//g.Use(middleware) // set a global middleware

	PackageRoutes[Response](g.Group("v1"), &A{})  //register the first service
	PackageRoutes[Response](g.Group("v2"), &A2{}) //register the second service

	group3 := g.Group("v3")
	//group3.Use(middleware) // set a middleware to this service
	PackageRoutes[Response](group3, &A3{}) //register the third service

	go func() {
		err := g.Run(fmt.Sprintf(":%d", 8080))
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Millisecond * 100)
	m.Run()
	os.Exit(0)
}

// Response defile you response struct type
// 自定义的返回结构
type Response struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

// A define a service
// 自定义的服务
type A struct{}

// GetRoutes defile routers
// 自定义的接口
func (a *A) GetRoutes() Routes[Response] {
	return Routes[Response]{
		{Method: GET, Path: "demo", Handle: a.HandleGetDemo},
		{Method: POST, Path: "demo", Handle: a.HandlePostDemo},
	}
}

func (a *A) HandleGetDemo(c *gin.Context) Response {
	//write logic here 在这里实现你的逻辑
	return Response{
		Code: 0,
		Desc: "OK",
		Data: map[string]any{"a": "a", "b": "b", "c": "c"},
	}
}

func (a *A) HandlePostDemo(c *gin.Context) Response {
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
		return Response{
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
	return Response{
		Code: 0,
		Desc: "OK",
		Data: data, // set response data here.
	}
}

type A2 struct{}

func (a *A2) GetRoutes() Routes[Response] {
	return Routes[Response]{
		{Method: GET, Path: "demo", Handle: a.HandleGetDemo},
		{Method: POST, Path: "demo", Handle: a.HandlePostDemo},
	}
}

func (a *A2) HandleGetDemo(c *gin.Context) Response {
	panic("not implemented")
}

func (a *A2) HandlePostDemo(c *gin.Context) Response {
	panic("not implemented")
}

type A3 struct{}

func (a *A3) GetRoutes() Routes[Response] {
	return Routes[Response]{
		{Method: GET, Path: "demo", Handle: func(c *gin.Context) Response {
			panic("not implemented")
		}},
		{Method: POST, Path: "demo", Handle: func(c *gin.Context) Response {
			panic("not implemented")
		}},
	}
}

func TestPackageDemo(t *testing.T) {
	{
		var result map[string]any
		resp, err := resty2.New().SetRetryCount(3).
			SetRetryWaitTime(time.Second * 2).
			R().
			SetResult(&result).
			SetQueryParams(map[string]string{}).Get("http://127.0.0.1:8080/v1/demo")
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode())
		t.Log(utilsregginroute.SoftNeatString(result))
	}
	t.Log("-")
	{
		var result map[string]any
		resp, err := resty2.New().SetRetryCount(3).
			SetRetryWaitTime(time.Second * 2).
			R().
			SetResult(&result).
			SetBody(map[string]any{"x": 1}).Post("http://127.0.0.1:8080/v1/demo")
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode())
		t.Log(utilsregginroute.SoftNeatString(result))
	}
}
