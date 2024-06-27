package reggin_test

import (
	"net/http/httptest"
	"testing"
	"time"

	resty2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/reggin/internal/demos/routers"
	"github.com/yyle88/reggin/internal/utils"
)

var caseServerUrxBase string

func TestMain(m *testing.M) {
	engine := routers.NewRouters()

	serverUt := httptest.NewServer(engine)
	defer serverUt.Close()

	caseServerUrxBase = serverUt.URL
	m.Run()
}

func TestDemo(t *testing.T) {
	var result map[string]any
	resp, err := resty2.New().SetRetryCount(3).
		SetRetryWaitTime(time.Second * 2).
		R().
		SetResult(&result).
		SetQueryParams(map[string]string{}).Get(caseServerUrxBase + "/v1/demo")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	t.Log(utils.SoftNeat(result))
}

func TestDemo2(t *testing.T) {
	var result map[string]any
	resp, err := resty2.New().SetRetryCount(3).
		SetRetryWaitTime(time.Second * 2).
		R().
		SetResult(&result).
		SetBody(map[string]any{"x": 1}).Post(caseServerUrxBase + "/v1/demo")
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	t.Log(utils.SoftNeat(result))
}
