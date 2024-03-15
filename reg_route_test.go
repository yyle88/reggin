package regginroute

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	resty2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/regginroute/utilsregginroute"
)

type A struct{}

func (a *A) GetRoutes() Routes[map[string]any] {
	return Routes[map[string]any]{
		{
			Method: GET,
			Path:   "demo",
			Handle: func(c *gin.Context) map[string]any {
				return map[string]any{"a": 1}
			},
		},
		{
			Method: POST,
			Path:   "demo",
			Handle: func(c *gin.Context) map[string]any {
				return map[string]any{"b": 2}
			},
		},
	}
}

func TestMain(m *testing.M) {
	g := gin.New()
	group := g.Group("v1")
	PackageRoutes[map[string]any](group, &A{})
	go func() {
		err := g.Run(fmt.Sprintf(":%d", 8080))
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Millisecond * 100)
	m.Run()
	os.Exit(0)
}

func TestPackage(t *testing.T) {
	{
		var result map[string]any
		resp, err := resty2.New().SetRetryCount(3).
			SetRetryWaitTime(time.Second * 2).
			R().
			SetResult(&result).
			SetQueryParams(map[string]string{}).Get("http://127.0.0.1:8080/v1/demo")
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode())
		t.Log(utilsregginroute.SoftNeatString(result))
	}
	t.Log("-")
	{
		var result map[string]any
		resp, err := resty2.New().SetRetryCount(3).
			SetRetryWaitTime(time.Second * 2).
			R().
			SetResult(&result).
			SetQueryParams(map[string]string{}).Post("http://127.0.0.1:8080/v1/demo")
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode())
		t.Log(utilsregginroute.SoftNeatString(result))
	}
}
