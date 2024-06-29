package reghan

import (
	"net/http"
	"testing"

	resty2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
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
