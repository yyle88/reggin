// 测试文件，使用 `TestMain` 和 `httptest.NewServer`。该测试文件通过模拟 `gin.Engine` 和 `gin.RouterGroup`，测试包内的核心函数是否能够正确注册路由。
package regginservice_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/reggin/regginservice"
)

// Mock implementation of RegGinServiceIFace interface
type mockGinService struct{}

func (s *mockGinService) RegEngine(engine *gin.Engine) {
	engine.GET("/test1/test2", func(c *gin.Context) {
		c.String(http.StatusOK, "gin service registered")
	})
}

// Mock implementation of GinRouteGroupIFace interface
type mockRouteGroup struct{}

func (r *mockRouteGroup) RegRoutes(group *gin.RouterGroup) {
	group.GET("/path2", func(c *gin.Context) {
		c.String(http.StatusOK, "route group registered")
	})
}

var (
	serverUt   *httptest.Server
	mockGinSrv regginservice.RegGinServiceIFace
	mockRoutes regginservice.GinRouteGroupIFace
)

func TestMain(m *testing.M) {
	// Set up Gin engine and mocks
	gin.SetMode(gin.TestMode)
	engine := gin.Default()

	// Initialize mock services
	mockGinSrv = &mockGinService{}
	mockRoutes = &mockRouteGroup{}

	// Register routes using regginservice functions
	regginservice.RegGinEngine(engine, mockGinSrv)
	regginservice.SetGinRouteGroup(engine, "/path1", mockRoutes)

	// Start a new test server for unit test
	serverUt = httptest.NewServer(engine)

	// Run the tests
	m.Run()

	// Cleanup after tests
	serverUt.Close()
}

// Test if the service route is correctly registered
func TestServiceRoute(t *testing.T) {
	resp, err := http.Get(serverUt.URL + "/test1/test2")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body := make([]byte, 1024)
	n, _ := resp.Body.Read(body)
	require.Contains(t, string(body[:n]), "gin service registered")
	resp.Body.Close()
}

// Test if the route group is correctly registered
func TestRouteGroup(t *testing.T) {
	resp, err := http.Get(serverUt.URL + "/path1/path2")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body := make([]byte, 1024)
	n, _ := resp.Body.Read(body)
	require.Contains(t, string(body[:n]), "route group registered")
}

/*
### 解释：

1. **
`mockGinService` 和 `mockRouteGroup`**：这是对 `regginservice.RegGinServiceIFace` 和 `regginservice.GinRouteGroupIFace` 的模拟实现，用来测试接口的具体行为。

`mockGinService` 实现了一个简单的 GET 路由 `/test`，而 `mockRouteGroup` 实现了一个路由组 `/group/group`。

2. **`TestMain`**：
- 初始化了 Gin 引擎，并使用 `regginservice.RegGinEngine` 和 `regginservice.SetGinRouteGroup` 注册路由。
- 使用 `httptest.NewServer` 启动一个 HTTP 测试服务器，它会监听测试中的请求。
- 在 `m.Run()` 运行所有测试之后，关闭测试服务器。

3. **测试函数**：
- **`TestServiceRoute`**：测试 `/test` 路由是否已通过 `RegGinEngine` 注册，并返回预期的响应。
- **`TestRouteGroup`**：测试 `/group/group` 路由是否已通过 `SetGinRouteGroup` 注册，并返回预期的响应。

4. **`assert`**：使用 `github.com/stretchr/testify/assert` 库来简化断言和测试条件。你可以通过 `go get github.com/stretchr/testify/assert` 安装该库。

通过这种方式，你可以确保 `regginservice` 包中的核心函数在实际应用中能按预期运行。
*/
