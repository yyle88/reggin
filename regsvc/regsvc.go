package regsvc

import "github.com/gin-gonic/gin"

type Service interface {
	RegEngine(engine *gin.Engine)
}

func RegSvc(engine *gin.Engine, svc Service) {
	svc.RegEngine(engine)
}

type SvcGrp interface {
	RegRoutes(group *gin.RouterGroup)
}

func RegSvcGroup(group *gin.RouterGroup, svcGrp SvcGrp) {
	svcGrp.RegRoutes(group)
}

func SetSvcGroup(engine *gin.Engine, relativePath string, svcGrp SvcGrp) {
	svcGrp.RegRoutes(engine.Group(relativePath))
}
