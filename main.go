package main

import (
	_ "byn/routers"
	_ "byn/sysinit"
	"byn/utils/auth"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println(auth.Casbin)
	beego.Run()
}
