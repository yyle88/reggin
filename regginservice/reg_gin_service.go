package regginservice

import "github.com/gin-gonic/gin"

type RegGinServiceIFace interface {
	RegEngine(engine *gin.Engine)
}

func RegGinEngine(engine *gin.Engine, service RegGinServiceIFace) {
	service.RegEngine(engine)
}

type GinRouteGroupIFace interface {
	RegRoutes(group *gin.RouterGroup)
}

func RegGinRouteGroup(group *gin.RouterGroup, routeGroup GinRouteGroupIFace) {
	routeGroup.RegRoutes(group)
}

func SetGinRouteGroup(engine *gin.Engine, relativePath string, routeGroup GinRouteGroupIFace) {
	routeGroup.RegRoutes(engine.Group(relativePath))
}
