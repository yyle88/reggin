package main

import (
	"fmt"

	"github.com/yyle88/done"
	"github.com/yyle88/reggin/internal/demos/routers"
)

func main() {
	engine := routers.NewRouters()
	done.Done(engine.Run(fmt.Sprintf(":%d", 8080)))
}
