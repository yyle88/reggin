package warpginhandle_test

import (
	"net/http"
	"net/http/httptest"
	"sort"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	restyv2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
	"github.com/yyle88/reggin/warpginhandle"
)

var testServerURL string

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
	engine.POST("/bbb", warpginhandle.P0(bbbHandle, warpResp[map[string]string]))
	//推荐使用这种写法这样在路由表里就能一眼看出调用的函数和返回的结果
	engine.POST("/ccc", warpginhandle.P1(cccHandle, parseArg[map[string]int], warpResp[map[string]string]))
	//推荐使用
	engine.POST("/ddd", warpginhandle.P1(dddHandle, warpginhandle.BIND[map[string]int], warpResp[map[string]string]))
	//使用普通JSON传递参数
	engine.POST("/eee", warpginhandle.PX(eeeHandle, warpResp[map[string]string]))

	//返回基本类型而非指针，测试返回基本类型，逻辑和前面的基本相同
	engine.POST("/fff", warpginhandle.P0(fffHandle, caseResp[string]))
	//返回基本类型而非指针
	engine.POST("/ggg", warpginhandle.PX(gggHandle, caseResp[int]))
	//返回基本类型而非指针
	engine.POST("/hhh", warpginhandle.P1(hhhHandle, parseArg[map[string]int], caseResp[bool]))
	//哦对返回数组也非指针
	engine.POST("/iii", warpginhandle.P1(iiiHandle, warpginhandle.BIND[map[string]int], caseResp[[]string]))
	//前面返回*map是不科学的，这里返回map相对好些，也是非指针的返回类型
	engine.POST("/jjj", warpginhandle.P1(jjjHandle, parseArg[map[string]int], caseResp[map[string]string]))
	//这里带 gin.Context 做参数的那种处理函数的逻辑
	engine.POST("/kkk", warpginhandle.C1(kkkHandle, warpginhandle.BIND[map[string]int], caseResp[map[string]string]))
	//这里带 gin.Context 做参数的，但这里处理函数的返回的指针类型
	engine.POST("/lll", warpginhandle.C1(lllHandle, parseArg[map[string]int], warpResp[map[string]string]))
	//这里替换成使用 form 取参数的逻辑，这是gin自带的
	engine.GET("/mmm", warpginhandle.C1(mmmHandle, warpginhandle.Q[mmmParam], warpResp[map[string]string]))
	//这里替换成使用 json 取参数的逻辑，假如想用 json 接收 QueryParams 就可以这样
	engine.GET("/nnn", warpginhandle.C1(nnnHandle, warpginhandle.QueryJson[nnnParam], warpResp[map[string]string]))
	//这里替换成使用 form 取参数的逻辑，假如想用 form 接收 QueryParams 就可以这样，但推荐使用上面的 gin 自带的通过 form 获取参数
	engine.GET("/ooo", warpginhandle.C1(oooHandle, warpginhandle.QueryForm[oooParam], warpResp[map[string]string]))

	engine.GET("/ppp", warpginhandle.R0(pppHandle, wrongResp))
	engine.GET("/qqq", warpginhandle.R1(qqqHandle, warpginhandle.QueryForm[qqqParam], wrongResp))
	engine.POST("/rrr", warpginhandle.RX(rrrHandle, wrongResp))
	engine.POST("/sss", warpginhandle.RX(sssHandle, wrongResp))

	serverUt := httptest.NewServer(engine)
	defer serverUt.Close()

	testServerURL = serverUt.URL
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

func parseArg[argType any](c *gin.Context) (arg *argType, err error) {
	var req argType
	if err = c.ShouldBindJSON(&req); err != nil {
		return nil, erero.WithMessage(err, "CAN NOT BIND REQ")
	}
	return &req, nil
}

func TestAaa(t *testing.T) {
	var res map[string]string
	response, err := restyv2.New().R().SetResult(&res).Post(testServerURL + "/aaa")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, "aaa", res["case"])
}

func TestBbb(t *testing.T) {
	var data map[string]string
	var res = respType{Data: &data}
	response, err := restyv2.New().R().SetResult(&res).Post(testServerURL + "/bbb")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, 0, res.Code)
	require.Equal(t, "bbb", data["case"])
}

func TestCcc(t *testing.T) {
	var data map[string]string
	var res = respType{Data: &data}
	response, err := restyv2.New().R().SetBody(map[string]int{
		"a": 100,
	}).SetResult(&res).Post(testServerURL + "/ccc")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, 0, res.Code)
	require.Equal(t, "ccc", data["case"])
	require.Equal(t, "100", data["a"])
}

func TestDdd(t *testing.T) {
	{
		var data map[string]string
		var res = respType{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{
			"a": 100,
		}).SetResult(&res).Post(testServerURL + "/ddd")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, 0, res.Code)
		require.Equal(t, map[string]string{"a": "100"}, data)
	}
	{
		var data map[string]string
		var res = respType{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{}).SetResult(&res).Post(testServerURL + "/ddd")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, -1, res.Code)
	}
}

func TestEee(t *testing.T) {
	var data map[string]string
	var res = respType{Data: &data}
	response, err := restyv2.New().R().SetBody(map[string]int{
		"a": 100,
	}).SetResult(&res).Post(testServerURL + "/eee")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, 0, res.Code)
	require.Equal(t, "eee", data["case"])
	require.Equal(t, "100", data["a"])
}

func TestFff(t *testing.T) {
	var data string
	var res = respType{Data: &data}
	response, err := restyv2.New().R().SetResult(&res).Post(testServerURL + "/fff")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, 0, res.Code)
	require.Equal(t, "abc", data)
}

func TestGgg(t *testing.T) {
	var data int
	var res = respType{Data: &data}
	response, err := restyv2.New().R().SetBody(map[string]int{
		"a": 100,
		"b": 200,
	}).SetResult(&res).Post(testServerURL + "/ggg")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode())
	t.Log(string(response.Body()))
	require.Equal(t, 0, res.Code)
	require.Equal(t, 300, data)
}

func TestHhh(t *testing.T) {
	{
		var data bool
		var res = respType{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
		}).SetResult(&res).Post(testServerURL + "/hhh")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, 0, res.Code)
		require.Equal(t, false, data)
	}
	{
		var data bool
		var res = respType{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
			"c": 300,
			"d": 400,
		}).SetResult(&res).Post(testServerURL + "/hhh")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, 0, res.Code)
		require.Equal(t, true, data)
	}
}

func TestIii(t *testing.T) {
	{
		var data []string
		var res = respType{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
		}).SetResult(&res).Post(testServerURL + "/iii")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, 0, res.Code)
		require.Equal(t, []string{"a", "b"}, data)
	}
	{
		var data []string
		var res = respType{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{}).SetResult(&res).Post(testServerURL + "/iii")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, -1, res.Code)
	}
}

func TestJjj(t *testing.T) {
	{
		var data map[string]string
		var res = respType{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{
			"a": 100,
			"b": 200,
		}).SetResult(&res).Post(testServerURL + "/jjj")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, 0, res.Code)
		require.Equal(t, map[string]string{"a": "100", "b": "200"}, data)
	}
	{
		var data map[string]string
		var res = respType{Data: &data}
		response, err := restyv2.New().R().SetBody(map[string]int{}).SetResult(&res).Post(testServerURL + "/jjj")
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode())
		t.Log(string(response.Body()))
		require.Equal(t, -1, res.Code)
	}
}
