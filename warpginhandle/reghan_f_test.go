package warpginhandle_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	restyv2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/reggin/warpginhandle"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func tttHandle(ctx *gin.Context) (map[string]string, error) {
	type tttParam struct {
		A int `form:"a"`
		B int `form:"b"`
		C int `form:"c"`
	}
	param := rese.P1(warpginhandle.QueryForm[tttParam](ctx))

	return map[string]string{
		"a": strconv.Itoa(param.A),
		"b": strconv.Itoa(param.B),
		"c": strconv.Itoa(param.C),
	}, nil
}

func TestTtt(t *testing.T) {
	var res map[string]string
	response, err := restyv2.New().R().SetQueryParams(map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}).SetResult(&res).Get(testServerURL + "/ttt")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
	require.Equal(t, "1", res["a"])
	require.Equal(t, "2", res["b"])
	require.Equal(t, "3", res["c"])
}

type uuuParam struct {
	A string `form:"a"`
}

func uuuHandle(ctx *gin.Context, arg *uuuParam) (map[string]int, error) {
	zaplog.LOG.Debug("param", zap.String("a", arg.A))
	res, err := strconv.Atoi(arg.A)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return map[string]int{
		"a": res,
	}, nil
}

func TestUuu_200(t *testing.T) {
	var res map[string]int
	response, err := restyv2.New().R().SetQueryParams(map[string]string{
		"a": "123",
	}).SetResult(&res).Get(testServerURL + "/uuu")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
	require.Equal(t, 123, res["a"])
}

func TestUuu_400(t *testing.T) {
	var res map[string]int
	response, err := restyv2.New().R().SetQueryParams(map[string]string{
		"a": "xyz",
	}).SetResult(&res).Get(testServerURL + "/uuu")
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
}

type vvvParam struct {
	A string `json:"a"`
}

func vvvHandle(ctx *gin.Context, arg *vvvParam) (map[string]int, error) {
	zaplog.LOG.Debug("param", zap.String("a", arg.A))
	res, err := strconv.Atoi(arg.A)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return map[string]int{
		"a": res,
	}, nil
}

func TestVvv_200(t *testing.T) {
	var res map[string]int
	response, err := restyv2.New().R().SetBody(&vvvParam{A: "123"}).SetResult(&res).Post(testServerURL + "/vvv")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
	require.Equal(t, 123, res["a"])
}

func TestVvv_400(t *testing.T) {
	var res map[string]int
	response, err := restyv2.New().R().SetBody(&vvvParam{A: "xyz"}).SetResult(&res).Post(testServerURL + "/vvv")
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
}

type wwwParam struct {
	A int `json:"a"`
}

func wwwHandle(ctx *gin.Context, arg *wwwParam) (map[string]string, error) {
	zaplog.LOG.Debug("param", zap.Int("a", arg.A))
	return map[string]string{
		"a": strconv.Itoa(arg.A),
	}, nil
}

func TestWww(t *testing.T) {
	var res map[string]string
	response, err := restyv2.New().R().SetBody(&wwwParam{A: 123}).SetResult(&res).Post(testServerURL + "/www")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
	require.Equal(t, "123", res["a"])
}
