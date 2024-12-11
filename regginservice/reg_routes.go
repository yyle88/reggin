package regginservice

import "github.com/gin-gonic/gin"

type EndpointHandler interface {
	RegisterRoutes(engine *gin.Engine)
}

func AddEndpoints(engine *gin.Engine, service EndpointHandler) {
	service.RegisterRoutes(engine)
}

func SetupService(engine *gin.Engine, service EndpointHandler) {
	service.RegisterRoutes(engine)
}

type RouteGroupHandler interface {
	RegisterRoutes(group *gin.RouterGroup)
}

func AddRouteGroup(group *gin.RouterGroup, routeGroup RouteGroupHandler) {
	routeGroup.RegisterRoutes(group)
}

func SetRouteGroup(engine *gin.Engine, relativePath string, routeGroup RouteGroupHandler) {
	routeGroup.RegisterRoutes(engine.Group(relativePath))
}

type IRouterHandler interface {
	RegisterRoutes(router gin.IRouter)
}

func AddRoutes(router gin.IRouter, handler IRouterHandler) {
	handler.RegisterRoutes(router)
}

func SetRoutes(router gin.IRouter, relativePath string, handler IRouterHandler) {
	handler.RegisterRoutes(router.Group(relativePath))
}
