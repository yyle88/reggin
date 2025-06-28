package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/reggin"
	"github.com/yyle88/reggin/internal/demos/demo3x/internal/message"
	"github.com/yyle88/reggin/internal/demos/demo3x/internal/service"
)

func SetRouters(engine gin.IRouter) {
	reggin.PackageRoutes[message.Response](engine.Group("v1"), &service.A1{}) //register the first service
	reggin.PackageRoutes[message.Response](engine.Group("v2"), &service.A2{}) //register the second service

	group3 := engine.Group("v3")
	//group3.Use(middleware) // set a middleware only to this service
	//group3.Use(middleware) // set a middleware only to this service
	reggin.PackageRoutes[message.Response](group3, &service.A3{}) //register the third service
	// if you think it can not meet your needs, you can also write your own func
	group3.GET("example", func(c *gin.Context) {
		//write logic
		c.JSON(http.StatusOK, message.Response{})
	})
	group3.POST("example", func(c *gin.Context) {
		//write logic
		c.JSON(http.StatusOK, message.Response{})
	})
}
