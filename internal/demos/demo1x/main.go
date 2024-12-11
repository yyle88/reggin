package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/must"
	"github.com/yyle88/reggin/internal/demos/demo1x/routers1x"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Recovery()) // set a global middleware
	//engine.Use(middleware) // set a global middleware
	//engine.Use(middleware) // set a global middleware

	routers1x.SetRouters(engine)

	must.Done(engine.Run(":8080"))
}
