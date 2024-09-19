package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/reggin"
	"github.com/yyle88/reggin/internal/demos/reggin_demo1x/message"
	service2 "github.com/yyle88/reggin/internal/demos/reggin_demo1x/service"
)

func NewGinEngineWithRouters() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery()) // set a global middleware
	//g.Use(middleware) // set a global middleware
	//g.Use(middleware) // set a global middleware

	reggin.PackageRoutes[message.Response](g.Group("v1"), &service2.A1{}) //register the first service
	reggin.PackageRoutes[message.Response](g.Group("v2"), &service2.A2{}) //register the second service

	group3 := g.Group("v3")
	//group3.Use(middleware) // set a middleware only to this service
	//group3.Use(middleware) // set a middleware only to this service
	reggin.PackageRoutes[message.Response](group3, &service2.A3{}) //register the third service
	// if you think it can not meet your needs, you can also write your own func
	group3.GET("example", func(c *gin.Context) {
		//write logic
		c.JSON(http.StatusOK, message.Response{})
	})
	group3.POST("example", func(c *gin.Context) {
		//write logic
		c.JSON(http.StatusOK, message.Response{})
	})
	return g
}
