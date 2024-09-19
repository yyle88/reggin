package main

import (
	"fmt"

	"github.com/yyle88/done"
	"github.com/yyle88/reggin/internal/demos/reggin_demo1x/routers"
)

func main() {
	engine := routers.NewGinEngineWithRouters()
	done.Done(engine.Run(fmt.Sprintf(":%d", 8080)))
}
