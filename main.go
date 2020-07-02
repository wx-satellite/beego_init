package main

import (
	_ "byn/routers"
	_ "byn/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
