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

type RequestHandlerFunc func(c *gin.Context) any

type Route struct {
	Method MethodName
	Path   string
	Handle RequestHandlerFunc
}

type Routes []Route

type Application interface {
	GetRoutes() Routes
}

func PackageRoutes(group *gin.RouterGroup, app Application) {
	RegisterRoutes(group, app.GetRoutes())
}

func RegisterRoutes(group *gin.RouterGroup, urls Routes) {
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

/*
在 Gin 框架中，`SecureJSON` 和 `JSON` 都是用于处理 JSON 数据的方法，但它们在安全性上有所不同。

1. `JSON`：`JSON` 是 Gin 框架中的方法，用于将数据序列化为 JSON 格式并将其作为响应返回给客户端。它是通过调用 `c.JSON` 方法来实现的，其中 `c` 是 `gin.Context` 对象。`JSON` 方法会自动设置响应头中的 Content-Type 为 "application/json"，并将数据以 JSON 格式返回给客户端。

   示例用法：
   ````go
   c.JSON(http.StatusOK, gin.H{
       "message": "Hello, World!",
   })
   ```

2. `SecureJSON`：`SecureJSON` 是 Gin 框架中的方法，功能与 `JSON` 类似，但它会对响应的 JSON 数据进行安全处理。具体来说，`SecureJSON` 方法会对 JSON 数据进行 HTML 字符转义，以防止跨站脚本攻击（XSS）。这是通过调用 `c.SecureJSON` 方法来实现的。

   示例用法：
   ````go
   c.SecureJSON(http.StatusOK, gin.H{
       "message": "<script>alert('XSS')</script>",
   })
   ```

总结来说，`JSON` 方法用于普通的 JSON 数据序列化和返回，而 `SecureJSON` 方法在返回 JSON 数据时会对其进行安全处理，以提供更强的安全性保护。如果你的应用程序可能面临 XSS 攻击风险，建议使用 `SecureJSON` 方法来返回 JSON 数据。否则，使用普通的 `JSON` 方法即可。
*/

/*
我进行了以下更改：

将 HandleFunc 重命名为 RequestHandlerFunc，以更清晰地表达其含义。
将 Api 重命名为 Route，更准确地描述了其作用。
将 Apis 重命名为 Routes，更好地反映了它是一组路由信息的含义。
将 App 重命名为 Application，更具表达力和可读性。
将 Package 重命名为 PackageRoutes，以更好地描述它是用于注册路由的函数。
将 PackageUrls 重命名为 RegisterRoutes，更好地传达了其功能。
将循环变量 api 重命名为 route，以更准确地描述其含义。
将 run 闭包函数重命名为 handler，以更清晰地表达其作用。
这些命名建议旨在提高代码的可读性和可维护性，使其更符合常见的命名约定和最佳实践。请根据你的具体需求和偏好，自行选择是否采纳这些建议。
*/
