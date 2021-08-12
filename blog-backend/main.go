package main

import (
	_ "blog-backend/boot"
	_ "blog-backend/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
