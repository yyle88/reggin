package regginservice

import "github.com/gin-gonic/gin"

type Service interface {
	RegEngine(engine *gin.Engine)
}

func RegGinEngine(engine *gin.Engine, service Service) {
	service.RegEngine(engine)
}

type GinRouteGroup interface {
	RegRoutes(group *gin.RouterGroup)
}

func RegGinRouteGroup(group *gin.RouterGroup, routeGroup GinRouteGroup) {
	routeGroup.RegRoutes(group)
}

func SetGinRouteGroup(engine *gin.Engine, relativePath string, routeGroup GinRouteGroup) {
	routeGroup.RegRoutes(engine.Group(relativePath))
}
