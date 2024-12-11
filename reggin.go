package reggin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	DELETE Method = "DELETE"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	ANY    Method = "ANY"
)

// RequestHandlerFunc handles HTTP requests and returns a response of type RES.
// RequestHandlerFunc 是 http 路由的处理函数，返回 RES 类型。
type RequestHandlerFunc[RES any] func(c *gin.Context) RES

type Route[RES any] struct {
	Method Method                  // HTTP method
	Path   string                  // Route path
	Handle RequestHandlerFunc[RES] // Request handler
}

type Routes[RES any] []*Route[RES]

// Application defines an interface that returns routes for the app.
// Application 定义了一个接口，返回应用程序的路由。
type Application[RES any] interface {
	GetRoutes() Routes[RES]
}

// PackageRoutes registers the app's routes into a gin RouterGroup.
// PackageRoutes 将应用程序的路由注册到 gin RouterGroup 中。
func PackageRoutes[RES any](group *gin.RouterGroup, app Application[RES]) {
	RegisterRoutes(group, app.GetRoutes())
}

// RegisterRoutes registers routes to the provided gin RouterGroup.
// RegisterRoutes 将路由注册到提供的 gin RouterGroup 中。
func RegisterRoutes[RES any](group *gin.RouterGroup, routes Routes[RES]) {
	for idx := range routes {
		var route = routes[idx] // Avoid using idx directly in the loop, use a temporary variable

		run := func(ctx *gin.Context) {
			response := route.Handle(ctx)
			ctx.SecureJSON(http.StatusOK, response)
		}

		switch route.Method {
		case GET:
			group.GET(route.Path, run)
		case POST:
			group.POST(route.Path, run)
		case DELETE:
			group.DELETE(route.Path, run)
		case PUT:
			group.PUT(route.Path, run)
		case PATCH:
			group.PATCH(route.Path, run)
		case ANY:
			group.Any(route.Path, run)
		default:
			group.Handle(string(route.Method), route.Path, run)
		}
	}
}
