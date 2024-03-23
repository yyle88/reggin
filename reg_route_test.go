package regginroute

import (
	"fmt"
	"os"
	"testing"
	"time"

	resty2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/regginroute/utilsregginroute"
)

func TestMain(m *testing.M) {
	fmt.Println("please exec go run demo/main/main.go")
	m.Run()
	os.Exit(0)
}

func TestPackageDemo(t *testing.T) {
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
			SetBody(map[string]any{"x": 1}).Post("http://127.0.0.1:8080/v1/demo")
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode())
		t.Log(utilsregginroute.SoftNeatString(result))
	}
}
