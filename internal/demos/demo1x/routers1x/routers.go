package routers1x

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/reggin"
	"github.com/yyle88/reggin/internal/demos/demo1x/message1x"
	service2 "github.com/yyle88/reggin/internal/demos/demo1x/service1x"
)

func SetRouters(engine gin.IRouter) {
	reggin.PackageRoutes[message1x.Response](engine.Group("v1"), &service2.A1{}) //register the first service
	reggin.PackageRoutes[message1x.Response](engine.Group("v2"), &service2.A2{}) //register the second service

	group3 := engine.Group("v3")
	//group3.Use(middleware) // set a middleware only to this service
	//group3.Use(middleware) // set a middleware only to this service
	reggin.PackageRoutes[message1x.Response](group3, &service2.A3{}) //register the third service
	// if you think it can not meet your needs, you can also write your own func
	group3.GET("example", func(c *gin.Context) {
		//write logic
		c.JSON(http.StatusOK, message1x.Response{})
	})
	group3.POST("example", func(c *gin.Context) {
		//write logic
		c.JSON(http.StatusOK, message1x.Response{})
	})
}
