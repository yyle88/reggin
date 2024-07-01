package reghan

import (
	"net/http"
	"sort"
	"strconv"
	"testing"

	resty2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

func fffHandle() (string, error) {
	return "abc", nil
}

func gggHandle(arg *map[string]int) (int, error) {
	res := 0
	for _, v := range *arg {
		res += v
	}
	return res, nil
}

func hhhHandle(arg *map[string]int) (bool, error) {
	cnt := len(*arg)
	res := cnt > 3
	return res, nil
}

func iiiHandle(arg *map[string]int) ([]string, error) {
	if len(*arg) == 0 {
		return nil, erero.New("wrong")
	}
	res := make([]string, 0, len(*arg))
	for k := range *arg {
		res = append(res, k)
	}
	sort.Strings(res)
	return res, nil
}

func jjjHandle(arg *map[string]int) (map[string]string, error) {
	if len(*arg) == 0 {
		return nil, erero.New("wrong")
	}
	res := make(map[string]string, len(*arg))
	for k, v := range *arg {
		res[k] = strconv.Itoa(v)
	}
	return res, nil
}

func TestFff(t *testing.T) {
	var data string
	var res = ResponseType{Data: &data}
	response, err := resty2.New().R().SetResult(&res).Post(caseServerUrxBase + "/fff")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, "abc", data)
}

func TestGgg(t *testing.T) {
	var data int
	var res = ResponseType{Data: &data}
	response, err := resty2.New().R().SetBody(map[string]int{
		"a": 100,
		"b": 200,
	}).SetResult(&res).Post(caseServerUrxBase + "/ggg")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, 300, data)
}

func TestHhh(t *testing.T) {
	{
		var data bool
		var res = ResponseType{Data: &data}
		response, err := resty2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
		}).SetResult(&res).Post(caseServerUrxBase + "/hhh")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, false, data)
	}
	{
		var data bool
		var res = ResponseType{Data: &data}
		response, err := resty2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
			"c": 300,
			"d": 400,
		}).SetResult(&res).Post(caseServerUrxBase + "/hhh")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, true, data)
	}
}

func TestIii(t *testing.T) {
	{
		var data []string
		var res = ResponseType{Data: &data}
		response, err := resty2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
		}).SetResult(&res).Post(caseServerUrxBase + "/iii")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, []string{"a", "b"}, data)
	}
	{
		var data []string
		var res = ResponseType{Data: &data}
		response, err := resty2.New().R().SetBody(map[string]int{}).SetResult(&res).Post(caseServerUrxBase + "/iii")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, -1, res.Code)
	}
}

func TestJjj(t *testing.T) {
	{
		var data map[string]string
		var res = ResponseType{Data: &data}
		response, err := resty2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
		}).SetResult(&res).Post(caseServerUrxBase + "/jjj")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, map[string]string{"a": "100", "b": "200"}, data)
	}
	{
		var data map[string]string
		var res = ResponseType{Data: &data}
		response, err := resty2.New().R().SetBody(map[string]int{}).SetResult(&res).Post(caseServerUrxBase + "/jjj")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, -1, res.Code)
	}
}
