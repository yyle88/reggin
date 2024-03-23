package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/regginroute"
	"github.com/yyle88/regginroute/demo/message"
	"github.com/yyle88/regginroute/demo/service"
	"github.com/yyle88/regginroute/utilsregginroute"
)

func main() {
	g := gin.New()
	g.Use(gin.Recovery()) // set a global middleware
	//g.Use(middleware) // set a global middleware
	//g.Use(middleware) // set a global middleware

	regginroute.PackageRoutes[message.Response](g.Group("v1"), &service.A{})  //register the first service
	regginroute.PackageRoutes[message.Response](g.Group("v2"), &service.A2{}) //register the second service

	group3 := g.Group("v3")
	//group3.Use(middleware) // set a middleware only to this service
	//group3.Use(middleware) // set a middleware only to this service
	regginroute.PackageRoutes[message.Response](group3, &service.A3{}) //register the third service
	// if you think it can not meet your needs, you can also write your own func
	group3.GET("example", func(c *gin.Context) {
		//write logic
		c.JSON(http.StatusOK, message.Response{})
	})
	group3.POST("example", func(c *gin.Context) {
		//write logic
		c.JSON(http.StatusOK, message.Response{})
	})

	utilsregginroute.AssertDone(g.Run(fmt.Sprintf(":%d", 8080)))
}
