package reghan

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	resty2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
)

var caseServerUrxBase string

func TestMain(m *testing.M) {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.POST("/aaa", func(c *gin.Context) {
		res := map[string]string{
			"case": "aaa",
		}
		c.JSON(http.StatusOK, res)
	})
	//没有参数
	engine.POST("/bbb", Handle0p(bbbHandle, GinResponse[map[string]string]))
	//推荐使用这种写法这样在路由表里就能一眼看出调用的函数和返回的结果
	engine.POST("/ccc", Handle1p(cccHandle, parseArg[map[string]int], GinResponse[map[string]string]))
	//推荐使用
	engine.POST("/ddd", Handle1p(dddHandle, BindJson[map[string]int], GinResponse[map[string]string]))
	//使用普通JSON传递参数
	engine.POST("/eee", HandleXp(eeeHandle, GinResponse[map[string]string]))

	//测试返回基本类型，逻辑和前面的基本相同
	//返回基本类型而非指针
	engine.POST("/fff", Handle0a(fffHandle, NewResponse[string]))
	//返回基本类型而非指针
	engine.POST("/ggg", HandleXa(gggHandle, NewResponse[int]))
	//返回基本类型而非指针
	engine.POST("/hhh", Handle1a(hhhHandle, parseArg[map[string]int], NewResponse[bool]))
	//哦对返回数组也非指针
	engine.POST("/iii", Handle1a(iiiHandle, parseArg[map[string]int], NewResponse[[]string]))
	//前面返回*map是不科学的，这里返回map相对好些，也是非指针的返回类型
	engine.POST("/jjj", Handle1a(jjjHandle, parseArg[map[string]int], NewResponse[map[string]string]))

	//这里带 gin.Context 做参数的那种处理函数的逻辑
	engine.POST("/kkk", Handle1b(kkkHandle, parseArg[map[string]int], NewResponse[map[string]string]))
	//这里带 gin.Context 做参数的，但这里处理函数的返回的指针类型
	engine.POST("/lll", Handle1c(lllHandle, parseArg[map[string]int], GinResponse[map[string]string]))

	serverUt := httptest.NewServer(engine)
	defer serverUt.Close()

	caseServerUrxBase = serverUt.URL
	m.Run()
}

func bbbHandle() (*map[string]string, error) {
	res := map[string]string{
		"case": "bbb",
	}
	return &res, nil
}

func cccHandle(arg *map[string]int) (*map[string]string, error) {
	res := map[string]string{}
	for k, v := range *arg {
		res[k] = strconv.Itoa(v)
	}
	res["case"] = "ccc"
	return &res, nil
}

func dddHandle(arg *map[string]int) (*map[string]string, error) {
	if len(*arg) == 0 {
		return nil, erero.New("wrong")
	}
	res := make(map[string]string, len(*arg))
	for k, v := range *arg {
		res[k] = strconv.Itoa(v)
	}
	return &res, nil
}

func eeeHandle(arg *map[string]int) (*map[string]string, error) {
	res := map[string]string{}
	for k, v := range *arg {
		res[k] = strconv.Itoa(v)
	}
	res["case"] = "eee"
	return &res, nil
}

func parseArg[argType any](c *gin.Context) (arg *argType, err error) {
	var req argType
	if erx := c.ShouldBindJSON(&req); erx != nil {
		return nil, erero.WithMessage(erx, "CAN NOT BIND REQ")
	}
	return &req, nil
}

func TestAaa(t *testing.T) {
	var res map[string]string
	response, err := resty2.New().R().SetResult(&res).Post(caseServerUrxBase + "/aaa")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, "aaa", res["case"])
}

func TestBbb(t *testing.T) {
	var data map[string]string
	var res = ResponseType{Data: &data}
	response, err := resty2.New().R().SetResult(&res).Post(caseServerUrxBase + "/bbb")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, "bbb", data["case"])
}

func TestCcc(t *testing.T) {
	var data map[string]string
	var res = ResponseType{Data: &data}
	response, err := resty2.New().R().SetBody(map[string]int{
		"a": 100,
	}).SetResult(&res).Post(caseServerUrxBase + "/ccc")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, "ccc", data["case"])
	require.Equal(t, "100", data["a"])
}

func TestDdd(t *testing.T) {
	{
		var data map[string]string
		var res = ResponseType{Data: &data}
		response, err := resty2.New().R().SetBody(map[string]int{
			"a": 100,
		}).SetResult(&res).Post(caseServerUrxBase + "/ddd")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, map[string]string{"a": "100"}, data)
	}
	{
		var data map[string]string
		var res = ResponseType{Data: &data}
		response, err := resty2.New().R().SetBody(map[string]int{}).SetResult(&res).Post(caseServerUrxBase + "/ddd")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, -1, res.Code)
	}
}

func TestEee(t *testing.T) {
	var data map[string]string
	var res = ResponseType{Data: &data}
	response, err := resty2.New().R().SetBody(map[string]int{
		"a": 100,
	}).SetResult(&res).Post(caseServerUrxBase + "/eee")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, "eee", data["case"])
	require.Equal(t, "100", data["a"])
}
