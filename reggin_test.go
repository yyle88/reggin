package reggin_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	restyv2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/reggin"
)

var testServerURL string

func TestMain(m *testing.M) {
	must.Equals(reggin.GET, http.MethodGet)
	must.Equals(reggin.POST, http.MethodPost)
	must.Equals(reggin.DELETE, http.MethodDelete)
	must.Equals(reggin.PATCH, http.MethodPatch)
	must.Equals(reggin.PUT, http.MethodPut)

	engine := gin.New()

	engine.Use(gin.Recovery())
	reggin.PackageRoutes[Response](engine.Group("aa"), &ServiceA{msg: "this-is-service-a"})
	reggin.PackageRoutes[Response](engine.Group("bb"), &ServiceB{msg: "this-is-service-b"})

	serverUt := httptest.NewServer(engine)
	defer serverUt.Close()

	testServerURL = serverUt.URL
	m.Run()
}

// Response defile you response struct type
// 自定义的返回结构
type Response struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}

type ServiceA struct {
	msg string
}

func (a *ServiceA) GetRoutes() reggin.Routes[Response] {
	return reggin.Routes[Response]{
		{Method: reggin.GET, Path: "demo", Handle: a.HandleGetDemo},
	}
}

func (a *ServiceA) HandleGetDemo(c *gin.Context) Response {
	return Response{
		Code: 0,
		Desc: "OK",
		Data: map[string]any{"msg": a.msg, "a": "a", "b": "b", "c": "c"},
	}
}

type ServiceB struct {
	msg string
}

func (a *ServiceB) GetRoutes() reggin.Routes[Response] {
	return reggin.Routes[Response]{
		{Method: reggin.POST, Path: "demo", Handle: a.HandlePostDemo},
	}
}

func (a *ServiceB) HandlePostDemo(c *gin.Context) Response {
	type paramType struct {
		X int `json:"x"`
	}
	var param paramType
	must.Done(c.ShouldBindJSON(&param))
	return Response{
		Code: 0,
		Desc: "OK",
		Data: map[string]any{"msg": a.msg, "a": "a", "b": "b", "c": "c", "x": param.X},
	}
}

func TestDemo_Get(t *testing.T) {
	var result map[string]any
	resp, err := restyv2.New().SetRetryCount(3).
		SetRetryWaitTime(time.Second * 2).
		R().
		SetResult(&result).
		SetQueryParams(map[string]string{}).
		Get(testServerURL + "/aa/demo")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	t.Log(neatjsons.S(result))
}

func TestDemo_Post(t *testing.T) {
	var result map[string]any
	resp, err := restyv2.New().SetRetryCount(3).
		SetRetryWaitTime(time.Second * 2).
		R().
		SetResult(&result).
		SetBody(map[string]any{"x": 1}).
		Post(testServerURL + "/bb/demo")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	t.Log(neatjsons.S(result))
}
