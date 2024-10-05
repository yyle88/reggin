package reggin_test

import (
	"net/http/httptest"
	"testing"
	"time"

	restyv2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/reggin/internal/demos/reggin_demo1x/routers_demo1x"
)

var caseUrxBase string

func TestMain(m *testing.M) {
	engine := routers_demo1x.NewEngineWithHttpRoute()

	serverUt := httptest.NewServer(engine)
	defer serverUt.Close()

	caseUrxBase = serverUt.URL
	m.Run()
}

func TestDemo_Get(t *testing.T) {
	var result map[string]any
	resp, err := restyv2.New().SetRetryCount(3).
		SetRetryWaitTime(time.Second * 2).
		R().
		SetResult(&result).
		SetQueryParams(map[string]string{}).
		Get(caseUrxBase + "/v1/demo")
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
		Post(caseUrxBase + "/v1/demo")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	t.Log(neatjsons.S(result))
}
