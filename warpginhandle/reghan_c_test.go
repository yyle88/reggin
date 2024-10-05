package warpginhandle

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

func TestKkk(t *testing.T) {
	{
		var data map[string]string
		var res = ResponseExample{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
		}).SetResult(&res).Post(caseServerUrxBase + "/kkk")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, map[string]string{"a": "100", "b": "200"}, data)
	}
	{
		var data map[string]string
		var res = ResponseExample{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{}).SetResult(&res).Post(caseServerUrxBase + "/kkk")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, -1, res.Code)
	}
}

func TestLll(t *testing.T) {
	{
		var data map[string]string
		var res = ResponseExample{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
		}).SetResult(&res).Post(caseServerUrxBase + "/lll")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, map[string]string{"a": "100", "b": "200"}, data)
	}
	{
		var data map[string]string
		var res = ResponseExample{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{}).SetResult(&res).Post(caseServerUrxBase + "/lll")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, -1, res.Code)
	}
}
