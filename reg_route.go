package regginroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MethodName string

const (
	Any    MethodName = "Any"
	GET    MethodName = "GET"
	POST   MethodName = "POST"
	DELETE MethodName = "DELETE"
	PATCH  MethodName = "PATCH"
	PUT    MethodName = "PUT"
)

// RequestHandlerFunc 就是需要实现每个api的处理逻辑
// 这里使用范型的考虑是，避免返回值中出现不符合类型的，比如笔误返回其它类型(常见的是 return err 这种笔误)
// 这样就能在编码时确保所有的返回值都是我们定义的数据格式(通常认为同一组api的返回值，遵循相同的数据格式，比如数据data，错误码code，错误信息msg等等字段)
type RequestHandlerFunc[T any] func(c *gin.Context) T

type Route[T any] struct {
	Method MethodName
	Path   string
	Handle RequestHandlerFunc[T]
}

type Routes[T any] []*Route[T]

type Application[T any] interface {
	GetRoutes() Routes[T]
}

func PackageRoutes[T any](group *gin.RouterGroup, app Application[T]) {
	RegisterRoutes(group, app.GetRoutes())
}

func RegisterRoutes[T any](group *gin.RouterGroup, urls Routes[T]) {
	for idx := range urls {
		var route = urls[idx] //注意：这里不能使用循环变量idx或者其他的，而是要使用临时变量，除非是go高版本已修复这个问题

		run := func(ctx *gin.Context) {
			response := route.Handle(ctx)
			ctx.SecureJSON(http.StatusOK, response)
		}

		switch route.Method {
		case Any:
			group.Any(route.Path, run)
		case GET:
			group.GET(route.Path, run)
		case POST:
			group.POST(route.Path, run)
		case DELETE:
			group.DELETE(route.Path, run)
		case PATCH:
			group.PATCH(route.Path, run)
		case PUT:
			group.PUT(route.Path, run)
		default:
			group.Handle(string(route.Method), route.Path, run)
		}
	}
}
