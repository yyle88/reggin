package warpginhandle_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	restyv2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func kkkHandle(ctx *gin.Context, arg *map[string]int) (map[string]string, error) {
	if len(*arg) == 0 {
		return nil, erero.New("wrong")
	}
	res := make(map[string]string, len(*arg))
	for k, v := range *arg {
		res[k] = strconv.Itoa(v)
	}
	return res, nil
}

func lllHandle(ctx *gin.Context, arg *map[string]int) (*map[string]string, error) {
	if len(*arg) == 0 {
		return nil, erero.New("wrong")
	}
	res := make(map[string]string, len(*arg))
	for k, v := range *arg {
		res[k] = strconv.Itoa(v)
	}
	return &res, nil
}

type mmmParam struct {
	Ax string `form:"a"`
	Bx string `form:"b"`
}

func mmmHandle(ctx *gin.Context, arg *mmmParam) (*map[string]string, error) {
	return &map[string]string{
		"a": arg.Ax,
		"b": arg.Bx,
	}, nil
}

type nnnParam struct {
	Ax string `json:"a"`
	Bx string `json:"b"`
}

func nnnHandle(ctx *gin.Context, arg *nnnParam) (*map[string]string, error) {
	return &map[string]string{
		"a": arg.Ax,
		"b": arg.Bx,
	}, nil
}

type oooParam struct {
	Ax string `form:"a"`
	Bx string `form:"b"`
}

func oooHandle(ctx *gin.Context, arg *oooParam) (*map[string]string, error) {
	return &map[string]string{
		"a": arg.Ax,
		"b": arg.Bx,
	}, nil
}

func TestKkk(t *testing.T) {
	{
		var data map[string]string
		var res = ExampleResponse{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
		}).SetResult(&res).Post(testServerURL + "/kkk")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, 0, res.Code)
		require.Equal(t, map[string]string{"a": "100", "b": "200"}, data)
	}
	{
		var data map[string]string
		var res = ExampleResponse{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{}).SetResult(&res).Post(testServerURL + "/kkk")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, -1, res.Code)
	}
}

func TestLll(t *testing.T) {
	{
		var data map[string]string
		var res = ExampleResponse{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
		}).SetResult(&res).Post(testServerURL + "/lll")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, 0, res.Code)
		require.Equal(t, map[string]string{"a": "100", "b": "200"}, data)
	}
	{
		var data map[string]string
		var res = ExampleResponse{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{}).SetResult(&res).Post(testServerURL + "/lll")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, -1, res.Code)
	}
}

func TestMmm(t *testing.T) {
	var data map[string]string
	var res = ExampleResponse{Data: &data}
	response, err := restyv2.New().R().SetQueryParams(map[string]string{
		"a": "aaa",
		"b": "bbb",
	}).SetResult(&res).Get(testServerURL + "/mmm")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, 0, res.Code)
	require.Equal(t, map[string]string{"a": "aaa", "b": "bbb"}, data)
}

func TestNnn(t *testing.T) {
	var data map[string]string
	var res = ExampleResponse{Data: &data}
	response, err := restyv2.New().R().SetQueryParams(map[string]string{
		"a": "aaa",
		"b": "bbb",
	}).SetResult(&res).Get(testServerURL + "/nnn")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, 0, res.Code)
	require.Equal(t, map[string]string{"a": "aaa", "b": "bbb"}, data)
}

func TestOoo(t *testing.T) {
	var data map[string]string
	var res = ExampleResponse{Data: &data}
	response, err := restyv2.New().R().SetQueryParams(map[string]string{
		"a": "aaa",
		"b": "bbb",
	}).SetResult(&res).Get(testServerURL + "/ooo")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, 0, res.Code)
	require.Equal(t, map[string]string{"a": "aaa", "b": "bbb"}, data)
}
