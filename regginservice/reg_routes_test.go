package regginservice_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	restyv2 "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/reggin/regginservice"
)

var testServerURL string

func TestMain(m *testing.M) {
	engine := gin.Default()

	regginservice.AddEndpoints(engine, &mockService1{})
	regginservice.SetRouteGroup(engine, "/api", &mockService2{})
	regginservice.SetRoutes(engine, "/abc", &mockService3{})

	serverUt := httptest.NewServer(engine)
	defer serverUt.Close()

	testServerURL = serverUt.URL
	m.Run()
}

type mockService1 struct{}

func (s *mockService1) RegisterRoutes(router *gin.Engine) {
	router.GET("/test/service", func(c *gin.Context) {
		c.String(http.StatusOK, "Service endpoint hit")
	})
}

type mockService2 struct{}

func (r *mockService2) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("/test/route", func(c *gin.Context) {
		c.String(http.StatusOK, "Route group endpoint hit")
	})
}

type mockService3 struct{}

func (m *mockService3) RegisterRoutes(router gin.IRouter) {
	router.GET("/case/route", func(c *gin.Context) {
		c.String(http.StatusOK, "Route endpoint hit")
	})
}

func TestAddEndpoints(t *testing.T) {
	resp, err := restyv2.New().SetRetryCount(3).
		SetRetryWaitTime(time.Second * 2).
		R().
		Get(testServerURL + "/test/service")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())
	message := string(resp.Body())
	t.Log(message)
	require.Contains(t, message, "Service endpoint hit")
}

func TestSetRouteGroup(t *testing.T) {
	resp, err := restyv2.New().SetRetryCount(3).
		SetRetryWaitTime(time.Second * 2).
		R().
		Get(testServerURL + "/api/test/route")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())
	message := string(resp.Body())
	t.Log(message)
	require.Contains(t, message, "Route group endpoint hit")
}

func TestSetRoutes(t *testing.T) {
	resp, err := restyv2.New().SetRetryCount(3).
		SetRetryWaitTime(time.Second * 2).
		R().
		Get(testServerURL + "/abc/case/route")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())
	message := string(resp.Body())
	t.Log(message)
	require.Contains(t, message, "Route endpoint hit")
}
