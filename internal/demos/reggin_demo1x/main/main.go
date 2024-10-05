package main

import (
	"fmt"

	"github.com/yyle88/done"
	"github.com/yyle88/reggin/internal/demos/reggin_demo1x/routers_demo1x"
)

func main() {
	engine := routers_demo1x.NewEngineWithHttpRoute()
	done.Done(engine.Run(fmt.Sprintf(":%d", 8080)))
}
