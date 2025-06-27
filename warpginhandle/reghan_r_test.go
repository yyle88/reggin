package warpginhandle_test

import (
	"net/http"
	"strconv"
	"testing"

	restyv2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func pppHandle() (map[string]string, error) {
	return map[string]string{
		"a": "x",
		"b": "y",
		"c": "z",
	}, nil
}

func TestPpp(t *testing.T) {
	var res map[string]string
	response, err := restyv2.New().R().SetResult(&res).Get(testServerURL + "/ppp")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
	require.Equal(t, "x", res["a"])
	require.Equal(t, "y", res["b"])
	require.Equal(t, "z", res["c"])
}

type qqqParam struct {
	A string `form:"a"`
}

func qqqHandle(arg *qqqParam) (map[string]int, error) {
	zaplog.LOG.Debug("param", zap.String("a", arg.A))
	res, err := strconv.Atoi(arg.A)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return map[string]int{
		"a": res,
	}, nil
}

func TestQqq_200(t *testing.T) {
	var res map[string]int
	response, err := restyv2.New().R().SetQueryParams(map[string]string{
		"a": "123",
	}).SetResult(&res).Get(testServerURL + "/qqq")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
	require.Equal(t, 123, res["a"])
}

func TestQqq_400(t *testing.T) {
	var res map[string]int
	response, err := restyv2.New().R().SetQueryParams(map[string]string{
		"a": "xyz",
	}).SetResult(&res).Get(testServerURL + "/qqq")
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
}

type rrrParam struct {
	A string `json:"a"`
}

func rrrHandle(arg *rrrParam) (map[string]int, error) {
	zaplog.LOG.Debug("param", zap.String("a", arg.A))
	res, err := strconv.Atoi(arg.A)
	if err != nil {
		return nil, erero.Wro(err)
	}
	return map[string]int{
		"a": res,
	}, nil
}

func TestRrr_200(t *testing.T) {
	var res map[string]int
	response, err := restyv2.New().R().SetBody(&rrrParam{A: "123"}).SetResult(&res).Post(testServerURL + "/rrr")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
	require.Equal(t, 123, res["a"])
}

func TestRrr_400(t *testing.T) {
	var res map[string]int
	response, err := restyv2.New().R().SetBody(&rrrParam{A: "xyz"}).SetResult(&res).Post(testServerURL + "/rrr")
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
}

type sssParam struct {
	A int `json:"a"`
}

func sssHandle(arg *sssParam) (map[string]string, error) {
	zaplog.LOG.Debug("param", zap.Int("a", arg.A))
	return map[string]string{
		"a": strconv.Itoa(arg.A),
	}, nil
}

func TestSss(t *testing.T) {
	var res map[string]string
	response, err := restyv2.New().R().SetBody(&sssParam{A: 123}).SetResult(&res).Post(testServerURL + "/sss")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(neatjsons.SxB(response.Body()))
	require.Equal(t, "123", res["a"])
}
