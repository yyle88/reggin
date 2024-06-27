# reggin

reggin means register gin routes. 非常简单的gin路由注册器。 

在我们使用gin做简单的http服务时，由于gin的处理函数类型为 `type HandlerFunc func(*Context)` 这就很不方便，按照某网红发明家的说辞就是，"很容易把我们累S"，因此我简单的对其进行了封装，让你能返回确定的自定义消息结构。

Demo:
[demoMain 文件](/demo/main/main.go)
[Unittest 文件](/reg_route_test.go)

## 详细的解释

比如以下为一种自定义结构：
```go
package message

type Response struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data any    `json:"data"`
}
```
当你在注册路由时使用：
```
reggin.PackageRoutes[message.Response](g.Group("v2"), &service.A2{})
```
这时就能限制你的处理函数返回值，只能是自定义的`message.Response`结构。

这样就能确保你的代码不会写错，你既不会忘记在必要的地方返回`c.JSON(http.StatusOK, resp)` 也不会在 `if err != nil { return err }` 的时候忘记把它写为`c.JSON(http.StatusOK, Response{Code:-1, Desc:"wrong"})` （在绝大多数情况下开发者为避免这个情况就会把逻辑分为service层和server层，只在server层返回`c.JSON(...)`，但这样会限制自由发挥，让程序变得比较紧巴）。

这将很大程度上提高你的编码效率，至少我自己是觉得挺好用的。

当你注册完服务
```
reggin.PackageRoutes[message.Response](g.Group("v2"), &service.A2{})
```

你就可以这样把路由注册进服务里：
```
type A2 struct{}

func (a *A2) GetRoutes() reggin.Routes[message.Response] {
	return reggin.Routes[message.Response]{
		{Method: reggin.GET, Path: "demo", Handle: a.HandleGetDemo},
		{Method: reggin.POST, Path: "demo", Handle: a.HandlePostDemo},
	}
}
```
这将使得你的路由非常清晰，代码也非常整洁。

接下来会提示你，函数未定义，直接定义函数即可。
```
func (a *A2) HandleGetDemo(c *gin.Context) message.Response {
	panic("not implemented")
}

func (a *A2) HandlePostDemo(c *gin.Context) message.Response {
	panic("not implemented")
}
```
这里你可以看到，由于没有server和service的两层结构，这里的逻辑不需要与外界交互，因此这里面的消息就可以定义为局部的。

比如这样：
```
func (a *A) HandleSetDemo(c *gin.Context) message.Response {
	type requestType struct {
		X int
	}
	type responseType struct {
		Y int
	}
	var req requestType
	if err := c.ShouldBindJSON(&req); err != nil {
		return message.Response{
			Code: -1,
			Desc: "wrong param",
			Data: nil,
		}
	}
	//write some logic
	res := req.X * 2
	//set return value
	return message.Response{
		Code: 0,
		Desc: "OK",
		Data: responseType{Y: res},
	}
}
```
你将非常清楚的知道，你的接口需要的参数和返回类型，这将让你的代码整体上是高内聚的。当然我建议你在函数的最开始就定义参数和返回两个类型。

这样能避免你在定义完server和service以后还要定义个巨大的message包(或param/response包)，里面是各种消息的结构，让你在阅读代码时来回跳转，这是没有用的（"很容易把我们累S"）。

这样能让你的代码极简，清晰，具有非常好的可读性，而且局部定义的东西修改的影响范围也确定。

当然即使是你想分开写也行，这都是可以的。

## 其它
这不是个框架，而只是个包，只适用于gin框架下的接口开发，在任意的地方你可以选择用它，或者不用它，或者只用一半，都是可以的。

我不太擅长写英文文档，基本都是中文写完然后让机器帮我翻译的，因此代码注释中也是有不少中文的，这对于中文开发者友好些。

## 谢谢
有兴趣的可以试用。

希望大家给个星星。
