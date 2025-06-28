package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yyle88/must"
	"github.com/yyle88/reggin/internal/demos/demo3x/internal/routers"
)

func main() {
	engine := gin.Default()
	//engine.Use(middleware) // set a global middleware
	//engine.Use(middleware) // set a global middleware

	routers.SetRouters(engine)

	must.Done(engine.Run(":8080"))
}
